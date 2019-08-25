package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GRPCOptions for starting the server.
type GRPCOptions struct {
	ServeAddr string
	Register  func(*grpc.Server)
}

// RunGRPCServer runs any service.
func RunGRPCServer(ctx context.Context, opts GRPCOptions) {
	lis, err := net.Listen("tcp", opts.ServeAddr)
	if err != nil {
		log.Fatal(err)
	}

	StartGRPCServer(ctx, lis, opts.Register)
	log.Println("Shutting down gRPC server")
}

// StartGRPCServer starts a server for any service.
// Stops the server when context is canceled.
func StartGRPCServer(ctx context.Context, lis net.Listener, register func(*grpc.Server)) {
	s := grpc.NewServer()
	defer s.GracefulStop()

	register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			if err != grpc.ErrServerStopped {
				log.Fatal(err)
			}
		}
	}()

	<-ctx.Done()
}

// ClientOptions for gRPC client.
type ClientOptions struct {
	Ctx           context.Context
	Addr          string
	Dialer        func(context.Context, string) (net.Conn, error)
	ClientFactory func(*grpc.ClientConn) interface{}
}

// NewGRPCClient returns a gRPC test client for any service.
// Closes the connection when context is canceled.
func NewGRPCClient(opts ClientOptions) (interface{}, error) {
	conn, err := grpc.DialContext(
		opts.Ctx,
		opts.Addr,
		grpc.WithContextDialer(opts.Dialer),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	go func() {
		<-opts.Ctx.Done()
		if err := conn.Close(); err != nil {
			log.Println("conn.Close():", err)
		}
	}()

	client := opts.ClientFactory(conn)

	return client, nil
}
