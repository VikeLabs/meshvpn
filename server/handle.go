package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/vikelabs/meshvpn/common/messages"
	"github.com/vikelabs/meshvpn/common/util"
)

type recievedRawMessage struct {
	conn net.Conn
	buf  []byte
}

type recievedMessage struct {
	conn net.Conn
	t    byte
	body []byte
}

func handleConnection(conn net.Conn) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, conn)
	if err != nil {
		util.ErrPrintf(
			"error reading from connection to %v: %v",
			conn.RemoteAddr().String(), err,
		)
		return
	}

	raw := recievedRawMessage{
		conn: conn,
		buf:  buf.Bytes(),
	}
	err = handleMessage(raw)
	if err != nil {
		util.ErrPrintln("error handling message:", err)
		return
	}
}

func handleMessage(raw recievedRawMessage) error {
	msg := parseMessage(raw)

	handlers.m.Lock()
	handler := handlers.h[msg.t]
	handlers.m.Unlock()

	return handler(msg)
}

func parseMessage(raw recievedRawMessage) recievedMessage {
	return recievedMessage{
		conn: raw.conn,
		t:    raw.buf[0],
		body: raw.buf[1:],
	}
}

type messageHandler func(msg recievedMessage) error
type syncHanders struct {
	h map[byte]messageHandler
	m sync.RWMutex
}

var handlers = syncHanders{
	h: map[byte]messageHandler{
		messages.PingType: handlePing,
		messages.PongType: handlePong,
	},
}

func handlePing(msg recievedMessage) error {
	resp := fmt.Sprint("Ping from " + msg.conn.RemoteAddr().String())
	_, err := msg.conn.Write([]byte(resp))
	return err
}

func handlePong(msg recievedMessage) error {
	resp := "Hello, " + string(msg.body) + "!"
	_, err := msg.conn.Write([]byte(resp))
	return err
}
