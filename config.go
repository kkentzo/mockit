package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
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

type Listeners struct {
	Listeners []*Listener `yaml:"listeners"`
}

type Listener struct {
	UriPath      string        `yaml:"uri_path"`
	Method       string        `yaml:"method"`
	ResponseCode int           `yaml:"response_code"`
	ResponseBody string        `yaml:"response_body"`
	Latency      time.Duration `yaml:"latency"`
	Headers      Headers       `yaml:"headers"`
}

type Config struct {
	file      string
	port      int
	listeners *Listeners
}

func ParseListenersFromFile(file string) (*Listeners, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	listeners := &Listeners{}
	err = yaml.Unmarshal(contents, listeners)
	if err != nil {
		return nil, err

	}
	// verify listeners
	for _, listener := range listeners.Listeners {
		if listener.UriPath == "" {
			return listeners,
				errors.New(fmt.Sprintf("Empty listener uri_path in file %s. Aborting.", file))
		}
	}
	return listeners, nil
}

func CreateConfig() *Config {
	config := &Config{}
	listener := &Listener{}
	flag.StringVar(&config.file, "config", "", "Path to configuration file")
	flag.IntVar(&config.port, "port", 9999, "HTTP Server Port")
	flag.BoolVar(&verbose, "verbose", true, "Activate logging")

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
		config.listeners, err = ParseListenersFromFile(config.file)
		if err != nil {
			log.Fatalf("Failed to parse file %s (%v)", config.file, err.Error())
		}
	}
	return config
}
