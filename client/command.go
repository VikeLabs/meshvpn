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

	reply, err := remoteCall(msg, body.Bytes())
	if err != nil {
		return errors.Wrap(err, "error ponging")
	}

	fmt.Printf("got response:\n%s\n", string(reply))
	return nil
}

func pongCommand(msg outgoingMessage) error {
	body := bytes.NewBuffer(nil)
	body.Write([]byte{messages.PongType})
	body.Write([]byte(pongName))

	reply, err := remoteCall(msg, body.Bytes())
	if err != nil {
		return errors.Wrap(err, "error ponging")
	}

	fmt.Printf("got response:\n%s\n", string(reply))
	return nil
}

func remoteCall(msg outgoingMessage, body []byte) ([]byte, error) {
	buf := bytes.NewBuffer(body)

	conn, err := net.Dial("tcp", msg.serverLoc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to server")
	}

	// send message
	conn.SetWriteDeadline(time.Now().Add(timeout))
	_, err = io.Copy(conn, buf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send message")
	}

	// get response
	conn.SetReadDeadline(time.Now().Add(timeout))
	_, err = io.Copy(buf, conn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get response")
	}
	conn.Close()

	return buf.Bytes(), nil
}
