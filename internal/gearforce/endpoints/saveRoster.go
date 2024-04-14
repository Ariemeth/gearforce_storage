package endpoints

import (
	"context"
	"encoding/json"
	"log"
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
			log.Println("Error saving roster: ", err)
			return saveRosterResponse{"", err}, nil
		}
		return saveRosterResponse{v.String(), nil}, nil
	}
}

func decodeSaveRosterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request saveRosterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Error decoding save roster request: ", err)
		return nil, errors.NewEndpointError(errors.ErrBadRosterFormat, ErrorStatusCode(errors.ErrBadRosterFormat))
	}
	return request, nil
}

type saveRosterRequest struct {
	Roster models.Roster `json:"roster"`
}

func (r *saveRosterRequest) UnmarshalJSON(data []byte) error {
	type Alias saveRosterRequest
	aux := &struct {
		Roster json.RawMessage `json:"roster"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var version models.ModelVersion
	if err := json.Unmarshal(aux.Roster, &version); err != nil {
		return err
	}

	switch version.Version {
	case 0, 1, 2:
		var roster models.RosterV2
		if err := json.Unmarshal(aux.Roster, &roster); err != nil {
			return err
		}
		r.Roster = roster
	default:
		var roster models.RosterV3
		if err := json.Unmarshal(aux.Roster, &roster); err != nil {
			return err
		}
		r.Roster = roster
	}

	return nil
}

type saveRosterResponse struct {
	Id  string `json:"id"`
	Err error  `json:"err,omitempty"`
}

func (r saveRosterResponse) Error() error {
	return r.Err
}
