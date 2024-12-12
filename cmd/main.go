package main

import (
	"github.com/sebastian-nunez/golang-key-value-db/core"
)

func main() {
	serverOpts := core.ServerOps{
		Port: 8081,
	}
	server := core.NewServer(serverOpts)
	server.Start()
}
