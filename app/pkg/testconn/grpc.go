package testconn

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// StartGRPCTestServer starts a test server for any service.
// Stops the server when context is canceled.
func StartGRPCTestServer(ctx context.Context, buf *BufNet, register func(*grpc.Server)) error {
	s := grpc.NewServer()
	register(s)

	go func() {
		<-ctx.Done()
		s.Stop()
	}()

	return s.Serve(buf.Listener)
}

// NewGRPCTestClient returns a gRPC test client for any service.
// Closes the connection when context is canceled.
func NewGRPCTestClient(ctx context.Context, buf *BufNet, newClient func(*grpc.ClientConn) interface{}) interface{} {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(buf.DialContext), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}

	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	return newClient(conn)
}
