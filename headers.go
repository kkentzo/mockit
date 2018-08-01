package main

import (
	"errors"
	"fmt"
	"strings"
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
			return errors.New(fmt.Sprintf("Invalid header specified: %s", keyVal))
		}
		(*h)[tokens[0]] = tokens[1]
	}
	return nil
}
