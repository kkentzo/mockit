package main

import (
	"flag"
	"fmt"
	"log"
)

type Config struct {
	file      string
	port      int
	endpoints *Endpoints
}

func CreateConfig() *Config {
	config := &Config{}
	endpoint := &Endpoint{}
	flag.StringVar(&config.file, "config", "", "Path to configuration file")
	flag.IntVar(&config.port, "port", 9999, "HTTP Server Port")

	flag.StringVar(&endpoint.UriPath, "uri", "/", "URI Path")
	flag.StringVar(&endpoint.Method, "method", "GET", "Request HTTP method")
	flag.IntVar(&endpoint.ResponseCode, "status", 200, "HTTP Response Status Code")
	flag.StringVar(&endpoint.ResponseBody, "body", "", "HTTP Response Body")
	flag.Var(&endpoint.Headers, "headers", "HTTP Response Headers (comma-separated)")
	flag.DurationVar(&endpoint.Latency, "latency", 0, "HTTP Response Latency")

	flag.Parse()

	if config.file == "" {
		config.endpoints = &Endpoints{Endpoints: []*Endpoint{endpoint}}
	} else {
		fmt.Printf("config file=%s\n", config.file)
		var err error
		config.endpoints, err = NewEndpointsFromFile(config.file)
		if err != nil {
			log.Fatalf("Failed to parse file %s (%v)", config.file, err.Error())
		}
	}
	return config
}
