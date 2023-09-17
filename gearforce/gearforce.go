package gearforce

import (
	"fmt"

	"github.com/Ariemeth/gearforce_storage/gearforce/endpoints"

	"github.com/gorilla/mux"
)

type gearForceService struct{}

func ConfigureRouteHandler(r *mux.Router, root string) {
	svc := gearForceService{}

	gfStoreHandler := endpoints.MakeHelloHTTPEndpointHandler(svc)

	r.Handle(root, gfStoreHandler).Methods("POST")
	r.Handle(fmt.Sprintf("%s/", root), gfStoreHandler).Methods("POST")
}

func (gearForceService) Hello(s string) (string, error) {

	return "hello", nil
}
