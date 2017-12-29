package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func mockHandler(w http.ResponseWriter, r *http.Request, config *Config) {
	time.Sleep(config.latency)
	// write the headers
	for key, val := range config.headers {
		w.Header().Set(key, val)
	}
	// write the response code
	w.WriteHeader(config.responseCode)
}

func main() {
	config := &Config{}
	ParseFromCommandLine(config)

	http.HandleFunc(config.uriPath, func(w http.ResponseWriter, r *http.Request) {
		mockHandler(w, r, config)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil))
}
