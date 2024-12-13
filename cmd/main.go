package main

import (
	"github.com/sebastian-nunez/golang-key-value-db/config"
	"github.com/sebastian-nunez/golang-key-value-db/core"
)

func main() {
	serverOpts := core.ServerOps{
		Port: config.Envs.Port,
	}
	server := core.NewServer(serverOpts)
	server.Start()
}
