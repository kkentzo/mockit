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
	// verify listeners
	for _, listener := range listeners.Listeners {
		if listener.UriPath == "" {
			return listeners,
				errors.New(fmt.Sprintf("Empty listener uri_path in file %s. Aborting.", file))
		}
	}
	// TODO: prevent mult regs of the same path
	return listeners, nil
}
