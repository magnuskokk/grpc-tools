package heartbeat

//go:generate mockgen -destination=mocks/service.go -package=mocks app/services/heartbeat ServiceServer
