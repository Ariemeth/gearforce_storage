package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ariemeth/gearforce_storage/internal/config"
	"github.com/Ariemeth/gearforce_storage/internal/gearforce"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Initializing service")

	log.Println("Loading environment variables")
	c, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatalf("Error: unable to load config: %v", err)
	}

	log.Println("Creating router")
	r := mux.NewRouter()

	gearforce.ConfigureRouteHandler(r.NewRoute().Subrouter(), "/gf", c.Database)

	r.HandleFunc("/healthz", healthHandler).Methods("GET")

	http.Handle("/", r)

	log.Println("Starting web server")
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", c.System.Port),
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
