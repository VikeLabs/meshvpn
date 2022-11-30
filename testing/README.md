# Testing

We use a network simulation for testing MeshVPN. This simulation uses network
namespaces to allow multiple Wireguard interfaces to interact on one machine,
and using Vagrant means it can run anywhere. The scripts use parts of the test
script included with `wireguard-go`, so you can check out the source material
[in the official repo](https://git.zx2c4.com/wireguard-go/tree/tests/netns.sh).

I still need to write proper docs, and since this is still a WIP, I'll just link
the most recent meeting notes and you can hopefully figure it out from there.

* https://hackmd.io/@malcolmseyd/rJEBB0RVd

I've copied the notes in their entirety below:

# MeshVPN Meeting 2021-03-29

The testing environment works! Let's take a look at how to use it.

## Getting set up

Installing Vagrant and Virtualbox is necessary to continue. Here are the steps you'll need to follow to get the testing environment set up.

* Create a new folder to work in
* Go into that folder
* Copy [__this code__](https://gist.github.com/malcolmseyd/b0a0e066c873e5dde130b61c3dd8104e) to a file named `Vagrantfile`
* Make a folder named `code`
* Run `vagrant up`

And you're good. From now on, these are the commands you'll need to know:

* Build/start the machine: `vagrant up`
* Login to the machine: `vagrant ssh`
* Shut the machine down: `vagrant halt`
* Destroy the machine: `vagrant destroy`

You'll want to use that last one if you get a new Vagrantfile or you just want to rebuild the machine from scratch.

## Using the testing environment

You'll log in as a regular user named `vagrant`. Most of the stuff requires root, so you'll want to run `sudo -s` to get a root shell. You can tell it's root because the prompt ends with `#` instead of `$`.

Next up, some commands:

* Test that the environment works at all: `netns.sh ./wireguard-go`
* Start the interfaces: `start.sh ./wireguard-go`
* Stop the interfaces: `stop.sh`

Cool. So those are some scripts that I wrote to help with setting up the Wireguard interfaces and the network namespaces. Note that `./wireguard-go` is the path to the Wireguard executable in `vagrant`'s home directory (`/home/vagrant`). If you're somewhere else, `/home/vagrant/wireguard-go` should work.

Ok, so what do the network interfaces look like? I've drawn a nice diagram to illustrate what's going on. Basically, wg1-3 were born in `ns0` and were moved to different network namespaces. All of their traffic still goes to and from `lo` in `ns0`, so all other interfaces appear to be at `lo`, that is, at `127.0.0.1`.

![image alt](https://i.imgur.com/niPIDQR.jpg")

That might be overwhelming, but don't worry, I've got some examples:

* List the network namespaces: `ip net`
* List the interfaces in `ns2`: `ip2 link`
* Run a program in `ns1`: `n1 <program>`
* List the Wireguard configurations for all interfaces: `n0 wg`
* Set a peer's endpoint: `n0 wg set wg3 peer /4UJX+UqbJ/EfSsjpo7dgLagMIHWr31iPHq/YggTAkM= endpoint 127.0.0.1:20000`

Basically, where # is some number, `ip#` runs `ip` in ns#, and `n#` runs any command in ns#. If you'd like to communicate with another interface, use `127.0.0.1` outside of Wireguard and `192.168.241.#` inside of Wireguard.

Here's an example workflow of what it might look like to test the program:

```shell
vagrant up
vagrant ssh
$ sudo -s
# start.sh ./wireguard-go
# n0 wg set wg1 peer "$pub2" endpoint 127.0.0.1:20000
# n0 wg set wg2 peer "$pub1" endpoint 127.0.0.1:10000
# cd ./code/server
# go build
# n1 ./server
#### in another terminal
vagrant ssh
$ sudo -s
# cd ./code/client
# n2 ./client
# exit
$ exit
#### back to the first terminal
<Ctrl-C>
# stop.sh
# exit
$ exit
vagrant halt
```

Hope that wasn't too confusing. Right now everthing connects directly, but  I'll improve the tooling in the future to support NATs.

## Tasks

* Get it to work on your machine
* Finish the code we discussed [a few weeks ago](/LNhNo_jER3SHrAAdb4KZpQ)

## Links

* [Vagrantfile](https://gist.github.com/malcolmseyd/b0a0e066c873e5dde130b61c3dd8104e)
* [Wireguard quick start](https://www.wireguard.com/quickstart/)
* [Guide for the `ip` command](https://baturin.org/docs/iproute2/)
* [The diagram (this link's mostly for me)](https://viewer.diagrams.net/?highlight=0000ff&edit=_blank&layers=1&nav=1&title=Untitled%20Diagram.drawio#R3VnZjtowFP2aPBbFdhZ4ZJu2UitVmofp9M0klyStiZExW7%2B%2BdnA2EgZGZSAzAgnfY18v5zr32MQi48Xus6DL%2BDsPgVnYDncWmVgYI%2BR56kcj%2BwPi%2Be4BiEQSmkYl8Jj8BQPaBl0nIaxqDSXnTCbLOhjwNIVA1jAqBN%2FWm805q4%2B6pBE0gMeAsib6lIQyNqtw7RL%2FAkkUm5HdvGJB87YGWMU05NsKRKYWGQvO5aG02I2Bae5yWg5%2BDydqi3kJSOUlDhP5dTgFkBz%2F%2BuF7EX2iQ%2FLJ9LKhbG3Wa2GPLpYWGaWzVfazyucv9zkngq%2FTEHS%2FtmqxjRMJj0sa6Nqt2gUKi%2BWCKQup4gaETBSfQ5ZEqcIk1w2osRjM1exH84SxMWdcZAOQ0IV%2B6Ch8JQX%2FA5WaPp4RtaXIyExcdQ67k4yggme1P4EvQIq9amIc8MAszWxNlNvbMtDIMVhcCTKyDUjN5oqKvssAqIKJwSvi0W%2BLB9MMzVQh0gXGc0iNUKD%2FF6MrkEkct0Zm8QBXyCQtXOK3onJwksow2eS0Iez3bPVBFU4r1S0eljuyyFB%2F3UkRG3HscNzNvYODj4JDmsHxW4LjvFVwtP%2FZxEPukHgo9OdBW%2BLxgj7M5m%2BTeNy75x1MzieebUTeRebx75x4iNOgBEJ1pjAmFzLmEU8pm5boqE5a2eYb17s2o%2Bo3SLk3ByS6lrxOJOwS%2BVO791xjPVdqJjvTc2bscyNV6604afO5Wle6ZVbpFw710UqZKU%2FhgDwkmqas%2FrB%2BveiXQ6o44msRwEsb05ApqYhAnlPO5h4RwKhMNvWJXP%2FxcS4SmwHuIa%2Ffww7qqUfpQXtdKiCNzuahSgaZENW76pryoHqqazti2TdVHvcC5cEfVXkcv3NHXtIk%2B5bpspIsy9R5Jl2iWrIsc2cH0qX3LtKld8lpA3fytOEeH6Xvfc%2FB%2Fmu1B19Pe3CXtcc5Oma33XpuLD6n7%2Fd1xo%2F06OT19GOqFELdU6nmZfSGKpWXX6VShdE9lRq8C5U6%2FQ9SRaVQJ1UKeR1TqXywy1UKXU%2BlUJdVCpHbqZQyyxcOWV3lrQ2Z%2FgM%3D)
