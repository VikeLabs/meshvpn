package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	flags := []cli.Flag{
		&cli.Uint64Flag{
			Name:     "port",
			Aliases:  []string{"p"},
			Usage:    "the port to listen on",
			Required: true,
		},
	}

	app := &cli.App{
		Name:   "server",
		Usage:  "Start a MeshVPN server",
		Action: run,
		Flags:  flags,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
