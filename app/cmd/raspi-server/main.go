package main

import (
	"app/api/raspi"
	"app/idl/raspi/raspiv1"
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
var grpcAddr = ":9001"
var httpAddr = ":8001"

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
				raspiv1.RegisterRaspiAPIServer(s, &raspi.API{})
			},
		})
	}()
	log.Println("Running raspi gRPC server at " + grpcAddr)

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunGatewayServer(ctx, &server.GatewayOptions{
			ServeAddr: httpAddr,
			GRPCAddr:  grpcAddr,
			DialOpts:  []grpc.DialOption{grpc.WithInsecure()},
			Register:  raspiv1.RegisterRaspiAPIHandlerFromEndpoint,
		})
	}()
	log.Printf("Running raspi HTTP server at " + httpAddr)

	<-sigs
	log.Println("Shutting down servers...")
}
