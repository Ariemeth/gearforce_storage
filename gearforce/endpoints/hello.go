package endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Ariemeth/gearforce_storage/gearforce/service"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHelloHTTPEndpointHandler(svc service.GearForceService) *httptransport.Server {
	handler := httptransport.NewServer(
		makeHelloEndpoint(svc),
		decodeHelloRequest,
		encodeResponse,
	)

	return handler
}

func makeHelloEndpoint(svc service.GearForceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(helloRequest)
		v, err := svc.Hello(req.Data)
		if err != nil {
			return helloResponse{v, err.Error()}, nil
		}
		return helloResponse{v, ""}, nil
	}
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request helloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type helloRequest struct {
	Data string `json:"data"`
}

type helloResponse struct {
	V   string `json:"results"`
	Err string `json:"err,omitempty"`
}
