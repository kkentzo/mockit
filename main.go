package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var verbose bool = true

func logRequest(r *http.Request, status int) {
	log.Printf("%s %s => %d", r.Method, r.RequestURI, status)
}

func mockHandler(w http.ResponseWriter, r *http.Request, listener *Listener) {
	// check the HTTP method
	if r.Method != strings.ToUpper(listener.method) {
		if verbose {
			logRequest(r, 404)
		}
		http.NotFound(w, r)
		return
	}
	if verbose {
		logRequest(r, listener.responseCode)
	}
	// enforce the latency
	time.Sleep(listener.latency)
	// write the headers
	for key, val := range listener.headers {
		w.Header().Set(key, val)
	}
	// write the response code
	w.WriteHeader(listener.responseCode)
}

func registerHandlers(listener *Listener) {
	http.HandleFunc(listener.uriPath, func(w http.ResponseWriter, r *http.Request) {
		mockHandler(w, r, listener)
	})
	if verbose {
		log.Printf("Listening: %s [method:%s|latency:%v]",
			listener.uriPath, listener.method, listener.latency)
	}
}

func main() {
	// do not print date/time info when logging
	log.SetFlags(0)

	config := &Config{}
	ParseFromCommandLine(config)

	if verbose {
		log.Printf("Server listens on port %d", config.port)
	}

	for _, listener := range config.listeners {
		registerHandlers(listener)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil))
}
