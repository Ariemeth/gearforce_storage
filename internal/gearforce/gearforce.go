package gearforce

import (
	"fmt"

	"github.com/Ariemeth/gearforce_storage/internal/gearforce/endpoints"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/models"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce/service"

	"github.com/gorilla/mux"
)

type GearForceService interface {
	service.Hello
	service.SaveRoster
}

type ServiceMiddleware func(GearForceService) GearForceService

func ConfigureRouteHandler(r *mux.Router, root string) {
	svc := gearForceService{}

	gfHelloHandler := endpoints.MakeHelloHTTPEndpointHandler(svc)

	r.Handle(fmt.Sprintf("%s/hello", root), gfHelloHandler).Methods("POST")
	r.Handle(fmt.Sprintf("%s/hello/", root), gfHelloHandler).Methods("POST")

	gfStoreHandler := endpoints.MakeStoreHTTPEndpointHandler(svc)
	r.Handle(fmt.Sprintf("%s/store", root), gfStoreHandler).Methods("POST")
	r.Handle(fmt.Sprintf("%s/store/", root), gfStoreHandler).Methods("POST")
}

type gearForceService struct{}

func (gearForceService) Hello(s string) (string, error) {

	return "hello", nil
}

func (gearForceService) Store(r models.Roster) (models.Roster, error) {

	return r, nil
}
