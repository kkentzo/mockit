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

type Config struct {
	uriPath      string
	method       string
	port         int
	responseCode int
	latency      time.Duration
	headers      Headers
}

func ParseFromCommandLine(config *Config) {
	flag.StringVar(&config.uriPath, "uri", "/", "URI Path")
	flag.StringVar(&config.method, "method", "GET", "Request HTTP method")
	flag.IntVar(&config.port, "port", 9999, "HTTP Server Port")
	flag.IntVar(&config.responseCode, "status", 200, "HTTP Response Status Code")
	flag.Var(&config.headers, "headers", "HTTP Response Headers (comma-separated)")
	flag.DurationVar(&config.latency, "latency", 0, "HTTP Response Latency")
	flag.Parse()
}
