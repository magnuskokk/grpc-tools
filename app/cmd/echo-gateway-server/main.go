package main

import (
	"app/idl/echo/echov1"
	"app/pkg/server"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

var serveAddr = ":8000"
var grpcAddr = "echo-grpc-server:9000"
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

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunGatewayServer(ctx, &server.GatewayOptions{
			ServeAddr: serveAddr,
			GRPCAddr:  grpcAddr,
			DialOpts:  []grpc.DialOption{grpc.WithInsecure()},
			Register:  echov1.RegisterEchoAPIHandlerFromEndpoint,
		})
	}()

	<-sigs
	log.Println("Shutting down servers...")
}
