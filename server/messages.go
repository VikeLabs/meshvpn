package main

import (
	"fmt"
	"sync"
)

type messageHandler func(msg recievedMessage) error

type handlerStruct struct {
	h map[byte]messageHandler
	m sync.RWMutex
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
