package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Endpoints struct {
	Endpoints []*Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	UriPath      string        `yaml:"uri_path"`
	Method       string        `yaml:"method"`
	ResponseCode int           `yaml:"response_code"`
	ResponseBody string        `yaml:"response_body"`
	Latency      time.Duration `yaml:"latency"`
	Headers      Headers       `yaml:"headers"`
}

func NewEndpointsFromFile(file string) (*Endpoints, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	endpoints := &Endpoints{}
	err = yaml.Unmarshal(contents, endpoints)
	if err != nil {
		return nil, err

	}
	return endpoints, endpoints.Validate()
}

func (endpoints *Endpoints) Validate() error {
	paths := make(map[string]bool)
	for _, endpoint := range endpoints.Endpoints {
		if endpoint.UriPath == "" {
			return errors.New("Empty uri_path in endpoint")
		}
		// make sure that no path is defined twice
		if _, ok := paths[endpoint.UriPath]; ok {
			return errors.New(fmt.Sprintf("Path %s already defined", endpoint.UriPath))
		} else {
			paths[endpoint.UriPath] = true
		}
	}
	return nil
}
