package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {
	if c.NArg() != 2 {
		return errors.New("error: command needs 2 arguments")
	}
	msg := outgoingMessage{
		serverLoc: c.Args().Get(0),
		msgType:   c.Args().Get(1),
	}

	return runCommand(msg)
}
