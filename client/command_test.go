package main

import (
	"bytes"
	"net"
	"testing"
)

func TestRemoteCall(t *testing.T) {
	// setup
	clientConn, serverConn := net.Pipe()

	expectedRequest := []byte("example request to the server")
	expectedResponse := []byte("example response from the server")

	actualRequest := make([]byte, 1024)

	// run this in the background, it'll block until the client uses the Conn
	go func() {
		n, err := serverConn.Read(actualRequest)
		if err != nil {
			t.Log("error reading from server pipe:", err)
			t.Fail()
		}
		actualRequest = actualRequest[:n]

		_, err = serverConn.Write(expectedResponse)
		if err != nil {
			t.Log("error writing to server pipe:", err)
			t.Fail()
		}
		serverConn.Close()
	}()

	// execution
	actualResponse, err := remoteCall(expectedRequest, clientConn)
	if err != nil {
		t.Log("unexpected error from remoteCall: ", err)
		t.Fail()
	}

	// checking
	if !bytes.Equal(expectedRequest, actualRequest) {
		t.Log("client didn't send correct bytes:",
			"\nexpected:", string(expectedRequest),
			"\nactual:  ", string(actualRequest),
		)
		t.Fail()
	}

	if !bytes.Equal(expectedResponse, actualResponse) {
		t.Log("client didn't recieve correct bytes:",
			"\nexpected:", string(expectedResponse),
			"\nactual:  ", string(actualResponse),
		)
		t.Fail()
	}
}
