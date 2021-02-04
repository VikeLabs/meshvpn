package main

import (
	"math"
	"net"
	"strconv"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {
	port := c.Uint64("port")
	if port == 0 || port > math.MaxUint16 {
		return errors.New("invalid port number, must be between 1 and 65535")
	}

	portStr := strconv.FormatUint(port, 10)
	address := net.JoinHostPort("", portStr)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Wrap(err, "error getting listener")
	}

	return listen(listener)
}

func listen(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return errors.Wrap(err, "error getting new connection")
		}

		go handleConnection(conn)
	}
}
