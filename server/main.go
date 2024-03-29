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
		&cli.StringFlag{
			Name:     "interface",
			Aliases:  []string{"i"},
			Usage:    "the Wireguard interface to use",
			Required: true,
		},
	}

	app := &cli.App{
		Name:            "server",
		Usage:           "Start a MeshVPN server",
		Action:          run,
		Flags:           flags,
		HideHelpCommand: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
