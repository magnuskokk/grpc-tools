package main

import (
	"app/api/echo"
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

var sigs chan os.Signal
var grpcAddr = ":9000"
var httpAddr = ":8000"

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
		server.RunGRPCServer(ctx, &server.GRPCOptions{
			ServeAddr: grpcAddr,
			Register: func(s *grpc.Server) {
				echov1.RegisterEchoAPIServer(s, &echo.API{})
			},
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunGatewayServer(ctx, &server.GatewayOptions{
			ServeAddr: httpAddr,
			GRPCAddr:  grpcAddr,
			DialOpts:  []grpc.DialOption{grpc.WithInsecure()},
			Register:  echov1.RegisterEchoAPIHandlerFromEndpoint,
		})
	}()

	<-sigs
	log.Println("Shutting down servers...")
}
