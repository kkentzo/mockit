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

func Handle(w http.ResponseWriter, r *http.Request, listener *Listener) {
	// check the HTTP method
	if r.Method != strings.ToUpper(listener.Method) {
		logRequest(r, 404)
		http.NotFound(w, r)
		return
	}
	logRequest(r, listener.ResponseCode)
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

func Register(listener *Listener, mux *http.ServeMux) {
	mux.HandleFunc(listener.UriPath, func(w http.ResponseWriter, r *http.Request) {
		Handle(w, r, listener)
	})
	log.Printf("Listening: %s [method:%s|latency:%v]",
		listener.UriPath, listener.Method, listener.Latency)
}
