package raspi

import (
	"app/idl/raspi/raspiv1"
	"context"
	"math/rand"
)

// API implements the generated RaspiAPIServer
type API struct {
}

// TempStream streams temperature data.
func (api *API) TempStream(req *raspiv1.TempStreamRequest, srv raspiv1.RaspiAPI_TempStreamServer) error {
	for {
		temp := &raspiv1.TempStreamResponse{
			Temp: &raspiv1.Temperature{
				Reading: int32(rand.Intn(100)),
			},
		}
		srv.Send(temp)
	}
}

// Radiator returns radiator reading
func (api *API) Radiator(ctx context.Context, req *raspiv1.RadiatorRequest) (*raspiv1.RadiatorResponse, error) {
	return &raspiv1.RadiatorResponse{
		Radi: &raspiv1.Radiator{
			Enabled: true,
			Level:   1,
		},
	}, nil
}

// SetRadiator updates and reads back radiator
func (api *API) SetRadiator(ctx context.Context, req *raspiv1.SetRadiatorRequest) (*raspiv1.SetRadiatorResponse, error) {
	return &raspiv1.SetRadiatorResponse{
		Radi: &raspiv1.Radiator{
			Enabled: false,
			Level:   0,
		},
	}, nil
}

// Status returns a combined reading.
func (api *API) Status(ctx context.Context, req *raspiv1.StatusRequest) (*raspiv1.StatusResponse, error) {
	return &raspiv1.StatusResponse{
		Status: &raspiv1.Status{
			Temperature: &raspiv1.Temperature{
				Reading: 1,
			},
			Radiator: &raspiv1.Radiator{
				Enabled: true,
				Level:   1,
			},
		},
	}, nil
}
