package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sebastian-nunez/golang-key-value-db/config"
	"github.com/sebastian-nunez/golang-key-value-db/core"
	"github.com/sebastian-nunez/golang-key-value-db/store"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signalCh
		fmt.Printf("\nReceived signal: '%s'. Gracefully shutting down...\n", sig)
		cancel()
	}()

	datastore := store.NewInMemoryStore()
	processor := core.NewCommandProcessor(datastore)

	serverOpts := core.TcpServerOps{
		Port: config.Envs.Port,
	}
	server := core.NewTcpServer(serverOpts, processor)
	go server.Start(ctx)

	<-ctx.Done()
	fmt.Println("Server has stopped.")
}
