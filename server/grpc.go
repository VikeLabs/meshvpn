package main

import (
	"context"

	"github.com/vikelabs/meshvpn/common/proto"
	"golang.zx2c4.com/wireguard/wgctrl"
)

// NewVPNServer creates a new VPN server
func NewVPNServer(wg *wgctrl.Client, wgName string) VPNServer {
	return VPNServer{
		proto.UnimplementedMeshVPNServer{},
		wg,
		wgName,
	}
}

// VPNServer is a protobuf server implementing MeshVPNServer
type VPNServer struct {
	proto.UnimplementedMeshVPNServer
	wg     *wgctrl.Client
	wgName string
}

func (VPNServer) mustEmbedUnimplementedMeshVPNServer() {}

// Ping is an rpc which simply returns.
func (s VPNServer) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingReply, error) {

	return &proto.PingReply{}, nil
}
