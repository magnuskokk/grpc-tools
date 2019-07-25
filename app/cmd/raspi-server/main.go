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

	register := func(s *grpc.Server) {
		raspiv1.RegisterRaspiAPIServer(s, &raspi.API{})
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.RunGRPCServer(ctx, grpcAddr, register); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Running raspi gRPC server at " + grpcAddr)

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.RunGatewayServer(ctx, &server.GatewayOptions{
			ServeAddr: httpAddr,
			GRPCAddr:  grpcAddr,
			DialOpts:  []grpc.DialOption{grpc.WithInsecure()},
			Register:  raspiv1.RegisterRaspiAPIHandlerFromEndpoint,
		})
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("Running raspi HTTP server at " + httpAddr)

	<-sigs
	log.Println("Shutting down servers...")
}
