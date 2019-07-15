package testing

import (
	"context"
	"net"

	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// BufNet implements a fake network for testing.
type BufNet struct {
	Listener *bufconn.Listener
}

// NewBufNet creates a new bufnet.
func NewBufNet() *BufNet {
	return &BufNet{
		Listener: bufconn.Listen(bufSize),
	}
}

// DialContext to satisfy function signature.
func (c *BufNet) DialContext(context.Context, string) (net.Conn, error) {
	return c.Listener.Dial()
}
