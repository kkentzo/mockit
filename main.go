package main

import (
	"log"
	"net/http"
)

func mockHandler(w http.ResponseWriter, r *http.Request, config *Config) {
}

func main() {
	config := &Config{}
	ParseFromCommandLine(config)

	http.HandleFunc(config.uriPath, func(w http.ResponseWriter, r *http.Request) {
		mockHandler(w, r, config)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
