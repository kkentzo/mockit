package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// do not print date/time info when logging
	log.SetFlags(0)

	config := CreateConfig()

	log.Printf("Server listens on port %d", config.port)

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.port),
		Handler: mux,
	}
	for _, listener := range config.listeners.Listeners {
		Register(listener, mux)
	}

	// start the server
	log.Fatal(srv.ListenAndServe())
}
