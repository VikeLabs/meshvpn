module github.com/vikelabs/meshvpn/server

go 1.15

replace github.com/vikelabs/meshvpn/common => ../common

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.3.0
	github.com/vikelabs/meshvpn/common v1.0.0
)
