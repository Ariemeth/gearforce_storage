package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	log.Println("Initializing service")

	//var gfSvc GearForceService
	gfSvc := gearForceService{}

	gfStoreHandler := httptransport.NewServer(
		makeUppercaseEndpoint(gfSvc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Handle("/", gfStoreHandler).Methods("POST")
	r.HandleFunc("/healthz", healthHandler).Methods("POST")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s1 := make(chan int)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
		s1 <- 0
	}()

	s2 := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(s2, os.Interrupt)

	// Block until we receive a termination signal.
	select {
	case <-s1:
	case <-s2:
	}

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy and running"))
}

type shortenRequest struct {
	Data string `json:"data"`
}

type shortenResponse struct {
	V   string `json:"short_url_code"`
	Err string `json:"err,omitempty"`
}

type GearForceService interface {
	Store(s string) (string, error)
}

type gearForceService struct{}

func (gearForceService) Store(s string) (string, error) {

	return "", nil
}

type ServiceMiddleware func(GearForceService) GearForceService

func makeUppercaseEndpoint(svc GearForceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(shortenRequest)
		v, err := svc.Store(req.Data)
		if err != nil {
			return shortenResponse{v, err.Error()}, nil
		}
		return shortenResponse{v, ""}, nil
	}
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
