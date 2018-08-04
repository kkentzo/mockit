package main

import (
	"fmt"
	"net"
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
	l, err := net.Listen("tcp", ":0")
	assert.Nil(t, err)

	Register(listener, mux)
	go http.Serve(l, mux)

	port := l.Addr().(*net.TCPAddr).Port
	path := fmt.Sprintf("http://localhost:%d/test", port)
	req, _ := http.NewRequest("GET", path, nil)
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
