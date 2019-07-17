package heartbeat

import context "context"

//go:generate mockgen -destination=mocks/service.go -package=mocks app/services/heartbeat ServiceServer

// Service for heartbeat server.
type Service struct {
}

// Ping responds with "pong".
func (s *Service) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return &PingReply{
		Message: req.Message,
	}, nil
}

// Stream loops and streams forever.
func (s *Service) Stream(req *StreamRequest, srv Service_StreamServer) error {
	for {
		packet := &StreamPacket{
			Id:   []byte("test"),
			Data: []byte("test"),
			Sum:  []byte("test"),
		}
		if err := srv.Send(packet); err != nil {
			return err
		}
	}
}

// Echo back with same request struct.
func (s *Service) Echo(ctx context.Context, req *EchoRequest) (*EchoRequest, error) {
	return req, nil
}
