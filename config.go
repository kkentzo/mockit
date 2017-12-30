package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

type Headers map[string]string

func (h *Headers) String() string {
	return fmt.Sprint(*h)
}

func (h *Headers) Set(value string) error {
	if *h == nil {
		*h = make(Headers)
	}
	for _, keyVal := range strings.Split(value, ",") {
		tokens := strings.SplitN(keyVal, ":", 2)
		if len(tokens) != 2 {
			log.Fatalf("Invalid header specified: %s", keyVal)
		}
		(*h)[tokens[0]] = tokens[1]
	}
	return nil
}

type Listener struct {
	uriPath      string
	method       string
	responseCode int
	latency      time.Duration
	headers      Headers
}

type Config struct {
	port      int
	listeners []*Listener
}

func ParseFromCommandLine(config *Config) {
	listener := &Listener{}
	flag.StringVar(&listener.uriPath, "uri", "/", "URI Path")
	flag.StringVar(&listener.method, "method", "GET", "Request HTTP method")
	flag.IntVar(&listener.responseCode, "status", 200, "HTTP Response Status Code")
	flag.Var(&listener.headers, "headers", "HTTP Response Headers (comma-separated)")
	flag.DurationVar(&listener.latency, "latency", 0, "HTTP Response Latency")

	flag.IntVar(&config.port, "port", 9999, "HTTP Server Port")
	flag.BoolVar(&verbose, "verbose", true, "Activate logging")

	flag.Parse()

	config.listeners = []*Listener{listener}
}
