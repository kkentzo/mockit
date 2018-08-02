package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

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

func NewListenersFromFile(file string) (*Listeners, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	listeners := &Listeners{}
	err = yaml.Unmarshal(contents, listeners)
	if err != nil {
		return nil, err

	}
	return listeners, listeners.Validate()
}

func (listeners *Listeners) Validate() error {
	paths := make(map[string]bool)
	for _, listener := range listeners.Listeners {
		if listener.UriPath == "" {
			return errors.New("Empty uri_path in listener")
		}
		// make sure that no path is defined twice
		if _, ok := paths[listener.UriPath]; ok {
			return errors.New(fmt.Sprintf("Path %s already defined", listener.UriPath))
		} else {
			paths[listener.UriPath] = true
		}
	}
	return nil
}
