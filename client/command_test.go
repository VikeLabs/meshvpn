package main

import (
	"bytes"
	"testing"
	"time"
)

type fakeDeadlineReadWriter struct {
	w  func([]byte) (int, error)
	r  func([]byte) (int, error)
	dw func(time.Time) error
	dr func(time.Time) error
}

func (f *fakeDeadlineReadWriter) Read(b []byte) (int, error) {
	return f.r(b)
}

func (f *fakeDeadlineReadWriter) Write(b []byte) (int, error) {
	return f.w(b)
}

func (f *fakeDeadlineReadWriter) SetReadDeadline(t time.Time) error {
	return f.dr(t)
}

func (f *fakeDeadlineReadWriter) SetWriteDeadline(t time.Time) error {
	return f.dw(t)
}

func TestRemoteCall(t *testing.T) {
	// setup
	noop := func(t time.Time) error { return nil }
	read := bytes.NewBuffer(nil)
	written := bytes.NewBuffer(nil)

	fakeConn := fakeDeadlineReadWriter{
		r:  read.Read,
		w:  written.Write,
		dr: noop,
		dw: noop,
	}

	body := []byte("test body data to send")

	replyString := "response from the server"
	read.WriteString(replyString)

	// execution
	reply, err := remoteCall(body, &fakeConn)
	if err != nil {
		t.Fatal("unexpected error from rpc: ", err)
	}

	// checking
	if !bytes.Equal(body, written.Bytes()) {
		t.Fatal("client didn't send correct bytes:",
			"\nexpected:", string(body),
			"\nactual:  ", string(written.Bytes()),
		)
	}

	if !bytes.Equal(reply, []byte(replyString)) {
		t.Fatal("client didn't recieve correct bytes:",
			"\nexpected:", replyString,
			"\nactual:  ", string(reply),
		)
	}
}
