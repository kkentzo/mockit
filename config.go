package main

import (
	"flag"
	"fmt"
	"log"
)

type Config struct {
	file      string
	port      int
	listeners *Listeners
}

func CreateConfig() *Config {
	config := &Config{}
	listener := &Listener{}
	flag.StringVar(&config.file, "config", "", "Path to configuration file")
	flag.IntVar(&config.port, "port", 9999, "HTTP Server Port")

	flag.StringVar(&listener.UriPath, "uri", "/", "URI Path")
	flag.StringVar(&listener.Method, "method", "GET", "Request HTTP method")
	flag.IntVar(&listener.ResponseCode, "status", 200, "HTTP Response Status Code")
	flag.StringVar(&listener.ResponseBody, "body", "", "HTTP Response Body")
	flag.Var(&listener.Headers, "headers", "HTTP Response Headers (comma-separated)")
	flag.DurationVar(&listener.Latency, "latency", 0, "HTTP Response Latency")

	flag.Parse()

	if config.file == "" {
		config.listeners = &Listeners{Listeners: []*Listener{listener}}
	} else {
		fmt.Printf("config file=%s\n", config.file)
		var err error
		config.listeners, err = NewListenersFromFile(config.file)
		if err != nil {
			log.Fatalf("Failed to parse file %s (%v)", config.file, err.Error())
		}
	}
	return config
}
