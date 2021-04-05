package main

import (
	"math"
	"net"
	"strconv"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/vikelabs/meshvpn/common/proto"
	"golang.zx2c4.com/wireguard/wgctrl"
	"google.golang.org/grpc"
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

	return listen(c, listener)
}

func listen(c *cli.Context, listener net.Listener) error {
	wg, err := wgctrl.New()
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	proto.RegisterMeshVPNServer(server, NewVPNServer(wg, c.String("interface")))
	return server.Serve(listener)
}
