package echo

//go:generate mockgen -destination=mocks/api.go -package=mocks app/idl/echo/echov1 EchoAPIServer

import (
	"app/idl/echo/echov1"
	"context"
)

// API implements the generated EchoAPIServer
type API struct {
}

// Echo returns the same request as a reply.
func (api *API) Echo(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {

	return &echov1.EchoResponse{
		Message: req.GetMessage(),
	}, nil
}
