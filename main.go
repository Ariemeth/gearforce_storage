package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Initializing service")

	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler).Methods("POST")
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
	URL string `json:"url"`
}

type shortenResponse struct {
	Shortened string `json:"short_url_code"`
	URL       string `json:"url"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var sr shortenRequest
	if err := decoder.Decode(&sr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: body does not contain url"))
	}

	key := createKey()

	resp := shortenResponse{Shortened: fmt.Sprintf("/%s", key), URL: sr.URL}
	json.NewEncoder(w).Encode(resp)
}

func createKey() string {

	return "a"
}
