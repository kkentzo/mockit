package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Register_WillSetupHandler(t *testing.T) {
	listener := &Listener{
		UriPath:      "/test",
		Method:       "GET",
		ResponseCode: 200,
		Headers:      Headers{"a": "b"},
	}
	mux := http.NewServeMux()
	srv := &http.Server{
		// TODO: how to select a free port for the http server?
		Addr:    ":32456",
		Handler: mux,
	}
	defer srv.Close()

	Register(listener, mux)
	go srv.ListenAndServe()

	req, _ := http.NewRequest("GET", "http://localhost:32456/test", nil)
	client := &http.Client{}
	res, err := client.Do(req)
	// TODO: getting connection refuse here
	assert.Nil(t, err)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		assert.Equal(t, 200, res.StatusCode)
	}
}
