package server

import (
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// RunGRPCServer is a wrapper function for easy running.
func RunGRPCServer(ctx context.Context, addr string, register func(s *grpc.Server)) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return StartGRPCServer(ctx, lis, register)
}

// StartGRPCServer starts a server for any service.
// Stops the server when context is canceled.
func StartGRPCServer(ctx context.Context, lis net.Listener, register func(*grpc.Server)) error {
	s := grpc.NewServer()
	register(s)

	errs := make(chan error)
	defer close(errs)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()

	go func() {
		defer wg.Done()
		if err := s.Serve(lis); err != nil {
			if err != grpc.ErrServerStopped {
				errs <- err
			}
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Shutting down gRPC server")
		s.GracefulStop()
		return nil

	case err := <-errs:
		return err
	}
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
func NewGRPCClient(opts *ClientOptions) (interface{}, error) {
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
