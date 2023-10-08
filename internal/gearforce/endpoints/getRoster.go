package endpoints

import (
	"context"
	"net/http"

	"github.com/Ariemeth/gearforce_storage/internal/gearforce/errors"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/models"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeGetRosterHTTPEndpointHandler(svc service.GetRoster) *httptransport.Server {
	handler := httptransport.NewServer(
		makeGetRosterEndpoint(svc),
		decodeGetRosterRequest,
		defaultEncodeResponse,
	)

	return handler
}

func makeGetRosterEndpoint(svc service.GetRoster) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRosterRequest)

		v, err := svc.GetRoster(req.ID)
		if err != nil {
			return getRosterResponse{v, err}, nil
		}
		return getRosterResponse{v, nil}, nil
	}
}

func decodeGetRosterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getRosterRequest
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		return request, errors.NewEndpointError(errors.ErrMissingId, ErrorStatusCode(errors.ErrMissingId))
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return request, errors.NewEndpointError(errors.ErrBadIdFormat, ErrorStatusCode(errors.ErrBadIdFormat))
	}

	request.ID = parsedId

	return request, nil
}

type getRosterRequest struct {
	ID uuid.UUID `json:"id"`
}

type getRosterResponse struct {
	Roster models.Roster `json:"roster"`
	Err    error         `json:"err,omitempty"`
}

func (r getRosterResponse) Error() error {
	return r.Err
}
