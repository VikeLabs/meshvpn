package main

import (
	"errors"
	"io"
	"net"

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

var handlers = handlerStruct{
	h: map[byte]messageHandler{
		messages.PingType: handlePing,
		messages.PongType: handlePong,
	},
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		util.ErrPrintf(
			"error reading from connection to %v: %v",
			conn.RemoteAddr().String(), err,
		)
		return
	}
	if n == 0 {
		util.ErrPrintln("error: response was length 0")
	}

	raw := recievedRawMessage{
		conn: conn,
		buf:  buf[:n],
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
