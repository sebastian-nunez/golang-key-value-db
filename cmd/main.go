package main

import (
	"github.com/sebastian-nunez/golang-key-value-db/core"
)

func main() {
	// TODO: create a `config` folder to handle variables coming directly from a `.env` file.
	serverOpts := core.ServerOps{
		Port: 3000,
	}
	server := core.NewServer(serverOpts)
	server.Start()
}
