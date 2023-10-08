package endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Ariemeth/gearforce_storage/internal/gearforce/errors"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/models"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeSaveRosterHTTPEndpointHandler(svc service.SaveRoster) *httptransport.Server {
	handler := httptransport.NewServer(
		makeSaveRosterEndpoint(svc),
		decodeSaveRosterRequest,
		defaultEncodeResponse,
	)

	return handler
}

func makeSaveRosterEndpoint(svc service.SaveRoster) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(saveRosterRequest)
		v, err := svc.SaveRoster(req.Roster)
		if err != nil {
			return saveRosterResponse{"", err}, nil
		}
		return saveRosterResponse{v.String(), nil}, nil
	}
}

func decodeSaveRosterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request saveRosterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.NewEndpointError(errors.ErrBadRosterFormat, ErrorStatusCode(errors.ErrBadRosterFormat))
	}
	return request, nil
}

type saveRosterRequest struct {
	Roster models.Roster `json:"roster"`
}

type saveRosterResponse struct {
	Id  string `json:"id"`
	Err error  `json:"err,omitempty"`
}

func (r saveRosterResponse) Error() error {
	return r.Err
}
