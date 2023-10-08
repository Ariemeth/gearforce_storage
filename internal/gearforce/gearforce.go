package gearforce

import (
	"encoding/json"
	"fmt"

	"github.com/Ariemeth/gearforce_storage/internal/gearforce/endpoints"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/errors"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/models"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/service"
	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

type GearForceService interface {
	service.Hello
	service.SaveRoster
	service.GetRoster
}

type ServiceMiddleware func(GearForceService) GearForceService

func ConfigureRouteHandler(r *mux.Router, root string) {
	svc := newGearForceService()

	gfHelloHandler := endpoints.MakeHelloHTTPEndpointHandler(svc)
	r.Handle(fmt.Sprintf("%s/hello", root), gfHelloHandler).Methods("POST")
	r.Handle(fmt.Sprintf("%s/hello/", root), gfHelloHandler).Methods("POST")

	gfStoreHandler := endpoints.MakeSaveRosterHTTPEndpointHandler(svc)
	r.Handle(fmt.Sprintf("%s/store", root), gfStoreHandler).Methods("POST")
	r.Handle(fmt.Sprintf("%s/store/", root), gfStoreHandler).Methods("POST")

	gfGetHandler := endpoints.MakeGetRosterHTTPEndpointHandler(svc)
	r.Handle(fmt.Sprintf("%s/{id}", root), gfGetHandler).Methods("GET")
	r.Handle(fmt.Sprintf("%s/{id}/", root), gfGetHandler).Methods("GET")
}

type gearForceService struct {
	Rosters map[uuid.UUID]models.Roster
}

func newGearForceService() GearForceService {
	svc := gearForceService{make(map[uuid.UUID]models.Roster)}

	return &svc
}

func (gearForceService) Hello(s string) (string, error) {

	return "hello", nil
}

// SaveRoster implements GearForceService.
func (g *gearForceService) SaveRoster(r models.Roster) (uuid.UUID, error) {
	key, _ := generatedID(r)
	g.Rosters[key] = r
	return key, nil
}

// GetRoster implements GearForceService.
func (g *gearForceService) GetRoster(id uuid.UUID) (models.Roster, error) {
	roster, exists := g.Rosters[id]
	if exists {
		return roster, nil
	}
	return models.Roster{}, errors.ErrIdNotFound
}

func generatedID(r models.Roster) (uuid.UUID, error) {
	u, err := json.Marshal(r)
	if err != nil {
		return uuid.New(), err
	}

	id := uuid.NewSHA1(uuid.MustParse("011f52f8-24e4-4bb9-a9e2-e13b7fcac716"), u)
	return id, nil
}
