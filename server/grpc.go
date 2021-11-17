package main

import (
	"context"
	"log"

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
	log.Printf("Recieved Ping:")
	return &proto.PingReply{}, nil
}

// ServerConnect is an rpc which returns the PublicKey and WireguardPort.
func (s VPNServer) ServerConnect(ctx context.Context, req *proto.ServerConnectRequest) (*proto.ServerConnectReply, error) {
	log.Printf("ClientPubkey:", string(req.GetClientPubkey()))

	conf, err := s.wg.Device(s.wgName)
	if err != nil {
		return nil, err
	}

	return &proto.ServerConnectReply{ServerPubkey: []byte(conf.PublicKey.String()), WireguardPort: int32(conf.ListenPort)}, nil
}
