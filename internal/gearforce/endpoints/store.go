package endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Ariemeth/gearforce_storage/internal/gearforce/models"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeStoreHTTPEndpointHandler(svc service.SaveRoster) *httptransport.Server {
	handler := httptransport.NewServer(
		makeStoreEndpoint(svc),
		decodeStoreRequest,
		defaultEncodeResponse,
	)

	return handler
}

func makeStoreEndpoint(svc service.SaveRoster) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(storeRequest)
		v, err := svc.Store(req.Roster)
		if err != nil {
			return storeResponse{v, err.Error()}, nil
		}
		return storeResponse{v, ""}, nil
	}
}

func decodeStoreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request storeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type storeRequest struct {
	Roster models.Roster `json:"roster"`
}

type storeResponse struct {
	Roster models.Roster `json:"roster"`
	Err    string        `json:"err,omitempty"`
}
