package main

import (
	"app/api/echo"
	echopb "app/generated/idl/echo"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
)

var sigs chan os.Signal

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

// fileserver conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileserver(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func main() {
	defer glog.Flush()

	// Start grpc server
	s := grpc.NewServer()
	echopb.RegisterEchoAPIServer(s, &echo.Service{})

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}

	defer lis.Close()
	defer s.Stop()

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
	err = echopb.RegisterEchoAPIHandlerFromEndpoint(ctx, mux, "localhost:9000", opts)
	if err != nil {
		glog.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Default().Handler)
	router.Handle("/*", mux)

	// provide swagger json from same server to enable swagger try out
	jsonDir := filepath.Join(os.Getenv("REPO_ROOT"), "swagger")
	fileserver(router, "/doc", http.Dir(jsonDir))

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
