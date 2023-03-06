package main

import (
	"fmt"
	"net"
	"net/rpc"

	goridgeRpc "github.com/roadrunner-server/goridge/v3/pkg/rpc"
)

type App struct{}

func (s *App) Hi(name string, r *string) error {
	*r = fmt.Sprintf("Hello, %s!", name)
	return nil
}

func main() {
	ln, err := net.Listen("unix", "/tmp/php-go-scylladb.sock")
	if err != nil {
		panic(err)
	}

	_ = rpc.Register(new(App))

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		_ = conn
		go rpc.ServeCodec(goridgeRpc.NewCodec(conn))
	}
}