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
	if r.Method != strings.ToUpper(listener.Method) {
		if verbose {
			logRequest(r, 404)
		}
		http.NotFound(w, r)
		return
	}
	if verbose {
		logRequest(r, listener.ResponseCode)
	}
	// enforce the latency
	time.Sleep(listener.Latency)
	// write the headers
	for key, val := range listener.Headers {
		w.Header().Set(key, val)
	}
	// write the response code
	w.WriteHeader(listener.ResponseCode)
	// write the response body
	w.Write([]byte(listener.ResponseBody))
}

func registerHandlers(listener *Listener) {
	http.HandleFunc(listener.UriPath, func(w http.ResponseWriter, r *http.Request) {
		mockHandler(w, r, listener)
	})
	if verbose {
		log.Printf("Listening: %s [method:%s|latency:%v]",
			listener.UriPath, listener.Method, listener.Latency)
	}
}

func main() {
	// do not print date/time info when logging
	log.SetFlags(0)

	config := CreateConfig()

	if verbose {
		log.Printf("Server listens on port %d", config.port)
	}

	for _, listener := range config.listeners.Listeners {
		registerHandlers(listener)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil))
}
