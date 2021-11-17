package main

import (
	"context"
	"fmt"
	"log"

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

	fmt.Println("Wireguard device info:")
	conf, err := wg.Device(wgDevName)
	if err != nil {
		return err
	}
	fmt.Println("PublicKey:", conf.PublicKey)
	fmt.Println("***********************")

	// Server connect
	r, err := client.ServerConnect(context.Background(), &proto.ServerConnectRequest{ClientPubkey: []byte(conf.PublicKey.String())})
	if err != nil {
		return err
	}
	log.Printf("ServerPubkey:", string(r.GetServerPubkey()))
	log.Printf("WireguardPort:", r.GetWireguardPort())

	return nil
}
