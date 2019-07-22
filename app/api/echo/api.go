package echo

import (
	. "app/generated/idl/echo"
	"context"
)

//go:generate mockgen -destination=mocks/api.go -package=mocks app/generated/idl/echo EchoAPIServer

// Service for echo server.
type Service struct {
}

// Echo returns the same request as a reply.
func (s *Service) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return &EchoResponse{
		Message: req.Message,
	}, nil
}
