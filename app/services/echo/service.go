package echo

import "context"

//go:generate mockgen -destination=mocks/service.go -package=mocks app/services/echo EchoServiceServer

// Service for echo server.
type Service struct {
}

// Echo returns the same request as a reply.
func (s *Service) Echo(ctx context.Context, req *EchoRequest) (*EchoRequest, error) {
	return req, nil
}
