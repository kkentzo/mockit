package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type Headers []string

func (h *Headers) String() string {
	return fmt.Sprint(*h)
}

func (h *Headers) Set(value string) error {
	for _, header := range strings.Split(value, ",") {
		*h = append(*h, header)
	}
	return nil
}

type Config struct {
	uriPath      string
	port         int
	responseCode int
	latency      time.Duration
	headers      Headers
}

func ParseFromCommandLine(config *Config) {
	flag.StringVar(&config.uriPath, "path", "/", "URI Path")
	flag.IntVar(&config.port, "port", 9999, "HTTP Server Port")
	flag.IntVar(&config.responseCode, "status", 200, "HTTP Response Status Code")
	flag.Var(&config.headers, "header", "HTTP Response Headers (comma-separated)")
	flag.DurationVar(&config.latency, "latency", 0, "HTTP Response Latency")
	flag.Parse()
}
