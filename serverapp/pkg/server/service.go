package server

//go:generate mockgen -destination=mocks/service.go -package=mocks serverapp/pkg/server HeartbeatServiceServer
