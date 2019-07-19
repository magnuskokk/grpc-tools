package ping

import "context"

// Service for ping server.
type Service struct {
}

// Ping function.
func (s *Service) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return &PingReply{
		Message: req.Message,
	}, nil
}
