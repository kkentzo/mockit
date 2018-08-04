package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func logRequest(r *http.Request, status int) {
	log.Printf("%s %s => %d", r.Method, r.RequestURI, status)
}

func Handle(w http.ResponseWriter, r *http.Request, endpoint *Endpoint) {
	// check the HTTP method
	if r.Method != strings.ToUpper(endpoint.Method) {
		logRequest(r, 404)
		http.NotFound(w, r)
		return
	}
	logRequest(r, endpoint.ResponseCode)
	// enforce the latency
	time.Sleep(endpoint.Latency)
	// write the headers
	for key, val := range endpoint.Headers {
		w.Header().Set(key, val)
	}
	// write the response code
	w.WriteHeader(endpoint.ResponseCode)
	// write the response body
	w.Write([]byte(endpoint.ResponseBody))
}

func Register(endpoint *Endpoint, mux *http.ServeMux) {
	mux.HandleFunc(endpoint.UriPath, func(w http.ResponseWriter, r *http.Request) {
		Handle(w, r, endpoint)
	})
	log.Printf("Endpoint: %s [method:%s|latency:%v]",
		endpoint.UriPath, endpoint.Method, endpoint.Latency)
}
