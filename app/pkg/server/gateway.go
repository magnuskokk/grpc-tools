package server

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

// GatewayOptions for HTTP gRPC gateway server.
type GatewayOptions struct {
	ServeAddr string
	GRPCAddr  string
	DialOpts  []grpc.DialOption
	Register  func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error
}

// NewGatewayMux dials and registers a mux for gateway.
func NewGatewayMux(ctx context.Context, opts *GatewayOptions) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	err := opts.Register(ctx, mux, opts.GRPCAddr, opts.DialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

// NewGatewayRouter wraps the mux in a chi.Router for modular mounting.
func NewGatewayRouter(ctx context.Context, opts *GatewayOptions) (chi.Router, error) {
	mux, err := NewGatewayMux(ctx, opts)
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Default().Handler)
	router.Handle("/*", mux)

	return router, nil
}

// NewGatewayTestServer returns a http test server for grpc gateway.
// The server is closed on context cancellation.
func NewGatewayTestServer(ctx context.Context, opts *GatewayOptions) (*httptest.Server, error) {
	mux, err := NewGatewayMux(ctx, opts)
	if err != nil {
		return nil, err
	}

	srv := httptest.NewServer(mux)

	go func() {
		<-ctx.Done()
		srv.Close()
	}()

	return srv, nil
}

// RunGatewayServer convienently dials to gRPC and runs the gateway.
func RunGatewayServer(ctx context.Context, opts *GatewayOptions) {
	router, err := NewGatewayRouter(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         opts.ServeAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Error shutting down http server:", err)
	}
}
