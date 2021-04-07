package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vikelabs/meshvpn/common/util"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "interface",
			Aliases:  []string{"i"},
			Usage:    "the Wireguard interface to use",
			Required: true,
		},
	}

	app := &cli.App{
		Name:            "client",
		Usage:           "Join a MeshVPN network",
		ArgsUsage:       "server:port",
		Flags:           flags,
		HideHelpCommand: true,
		Action:          run,
	}

	err := app.Run(os.Args)
	if err != nil {
		util.ErrPrintln(err)
	}
}
