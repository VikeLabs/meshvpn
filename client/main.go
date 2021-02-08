package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:            "client",
		Usage:           "Join a MeshVPN network",
		ArgsUsage:       "server:port messagetype",
		HideHelpCommand: true,
		Action:          run,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
