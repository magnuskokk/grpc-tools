package main

import (
	"app/pkg/metrics"
	"app/pkg/server"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var sigs chan os.Signal

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// gRPC endpoint
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunHTTPServer(ctx, server.HTTPOptions{
			ServeAddr: os.Getenv("HTTP_BIND_ADDR"),
			Handler:   metrics.NewRouter(),
		})
	}()

	fmt.Println("Running metrics server at", os.Getenv("HTTP_BIND_ADDR"))

	<-sigs
}
