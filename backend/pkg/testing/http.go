package testing

import (
	"context"
	"net/http/httptest"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// NewGatewayTestServer returns a http test server for grpc gateway.
// The server is closed on context cancellation.
func NewGatewayTestServer(ctx context.Context, buf *BufNet, register func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error) (*httptest.Server, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithContextDialer(buf.DialContext),
		grpc.WithInsecure(),
	}

	err := register(ctx, mux, "bufnet", opts)
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
