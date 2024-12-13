package main

import (
	"github.com/sebastian-nunez/golang-key-value-db/config"
	"github.com/sebastian-nunez/golang-key-value-db/core"
)

func main() {
	serverOpts := core.TcpServerOps{
		Port: config.Envs.Port,
	}
	server := core.NewTcpServer(serverOpts)
	server.Start()
}
