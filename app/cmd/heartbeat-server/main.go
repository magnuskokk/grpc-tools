package main

import (
	"app/services/heartbeat"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var sigs chan os.Signal

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	defer glog.Flush()

	// Start grpc server
	s := grpc.NewServer()
	heartbeat.RegisterServiceServer(s, &heartbeat.Service{})

	lis, err := net.Listen("tcp", ":9000")

	defer lis.Close()
	defer s.Stop()

	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			glog.Fatalf("server crashed: %v", err)
		}
	}()

	fmt.Println("Running gRPC server at localhost:9000")

	// Setup http gateway
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// TODO ServeMux opts
	// TODO use some other
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = heartbeat.RegisterServiceHandlerFromEndpoint(ctx, mux, "localhost:9000", opts)
	if err != nil {
		glog.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Handle("/*", mux)

	srv := &http.Server{
		Addr:         ":8081",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	defer srv.Shutdown(ctx)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				glog.Fatal(err)
			}
			log.Println("http server closed")
		}
	}()

	fmt.Println("Running HTTP server at localhost:8081")

	<-sigs
}
