package server

//go:generate mockgen -destination=mocks/service.go -package=mocks grpc-tools/pkg/server HeartbeatServiceServer
