# MeshVPN

MeshVPN is a mesh VPN control plane written in Golang. It uses Wireguard as a data plane. MeshVPN allows P2P connections between the members of the network, even if both members are behind a NAT. It uses a centralized control server to find existing peers and add new ones.

This project is in very early alpha and not all of the features mentions in this README may have been implemented yet.