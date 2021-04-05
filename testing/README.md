# Testing

We use a network simulation for testing MeshVPN. This simulation uses network
namespaces to allow multiple Wireguard interfaces to interact on one machine,
and using Vagrant means it can run anywhere. The scripts use parts of the test
script included with `wireguard-go`, so you can check out the source material
[in the official repo](https://git.zx2c4.com/wireguard-go/tree/tests/netns.sh).

I still need to write proper docs, and since this is still a WIP, I'll just link
the most recent meeting notes and you can hopefully figure it out from there.

* https://hackmd.io/@malcolmseyd/rJEBB0RVd