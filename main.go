package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func logRequest(r *http.Request, status int) {
	log.Printf("%s %s => %d", r.Method, r.RequestURI, status)
}

func mockHandler(w http.ResponseWriter, r *http.Request, config *Config) {
	// check the HTTP method
	if r.Method != strings.ToUpper(config.method) {
		if config.verbose {
			logRequest(r, 404)
		}
		http.NotFound(w, r)
		return
	}
	if config.verbose {
		logRequest(r, config.responseCode)
	}
	// enforce the latency
	time.Sleep(config.latency)
	// write the headers
	for key, val := range config.headers {
		w.Header().Set(key, val)
	}
	// write the response code
	w.WriteHeader(config.responseCode)
}

func main() {
	// do not print date/time info when logging
	log.SetFlags(0)

	config := &Config{}
	ParseFromCommandLine(config)

	http.HandleFunc(config.uriPath, func(w http.ResponseWriter, r *http.Request) {
		mockHandler(w, r, config)
	})

	if config.verbose {
		log.Printf("Listening on localhost:%d%s [method:%s|latency:%v]",
			config.port, config.uriPath, config.method, config.latency)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil))
}
