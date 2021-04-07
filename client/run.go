package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/vikelabs/meshvpn/common/proto"
	"github.com/vikelabs/meshvpn/common/util"
	"golang.zx2c4.com/wireguard/wgctrl"
	"google.golang.org/grpc"
)

func run(c *cli.Context) error {
	if c.NArg() != 1 {
		util.ErrPrintln("Error: expected 1 argument, but got ", c.NArg())
		cli.ShowAppHelpAndExit(c, 1)
	}

	serverLocation := c.Args().Get(0)
	wgDevName := c.String("interface")

	wg, err := wgctrl.New()
	if err != nil {
		return err
	}
	defer wg.Close()

	conn, err := grpc.Dial(serverLocation, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := proto.NewMeshVPNClient(conn)

	fmt.Println("Pinging...")
	_, err = client.Ping(context.Background(), &proto.PingRequest{})
	if err != nil {
		return err
	}
	fmt.Println("Ping successful!")

	fmt.Println("Wireguard device info:")
	fmt.Println(wg.Device(wgDevName))

	return nil
}
