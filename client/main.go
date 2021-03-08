package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vikelabs/meshvpn/common/util"
)

func main() {
	app := &cli.App{
		Name:            "client",
		Usage:           "Join a MeshVPN network",
		ArgsUsage:       "server:port wg_iface",
		HideHelpCommand: true,
		Action:          run,
	}

	err := app.Run(os.Args)
	if err != nil {
		util.ErrPrintln(err)
	}
}
