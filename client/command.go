package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/vikelabs/meshvpn/common/messages"
)

type outgoingMessage struct {
	serverLoc string
	msgType   string
}

const pongName = "William"
const timeout = time.Second * 3

func runCommand(msg outgoingMessage) error {
	switch msg.msgType {
	case "ping":
		return pingCommand(msg)
	case "pong":
		return pongCommand(msg)
	}

	return errors.New("invalid message type")
}

func pingCommand(msg outgoingMessage) error {
	body := bytes.NewBuffer(nil)
	body.Write([]byte{messages.PingType})

	conn, err := dialServer(msg)
	if err != nil {
		return err
	}

	reply, err := remoteCall(body.Bytes(), conn)
	if err != nil {
		return errors.Wrap(err, "error ponging")
	}

	conn.Close()

	fmt.Printf("got response:\n%s\n", string(reply))
	return nil
}

func pongCommand(msg outgoingMessage) error {
	body := bytes.NewBuffer(nil)
	body.Write([]byte{messages.PongType})
	body.Write([]byte(pongName))

	conn, err := dialServer(msg)
	if err != nil {
		return err
	}

	reply, err := remoteCall(body.Bytes(), conn)
	if err != nil {
		return errors.Wrap(err, "error ponging")
	}

	conn.Close()

	fmt.Printf("got response:\n%s\n", string(reply))
	return nil
}

type deadlineReadWriter interface {
	Read(b []byte) (int, error)
	Write(b []byte) (int, error)
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

func remoteCall(body []byte, rw deadlineReadWriter) ([]byte, error) {
	buf := bytes.NewBuffer(body)

	// send message
	rw.SetWriteDeadline(time.Now().Add(timeout))
	_, err := io.Copy(rw, buf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send message")
	}

	// get response
	rw.SetReadDeadline(time.Now().Add(timeout))
	_, err = io.Copy(buf, rw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get response")
	}
	return buf.Bytes(), nil
}

func dialServer(msg outgoingMessage) (net.Conn, error) {
	conn, err := net.Dial("tcp", msg.serverLoc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to server")
	}
	return conn, nil
}
