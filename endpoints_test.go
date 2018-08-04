package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TouchFile(t *testing.T, fname string) *os.File {
	f, err := os.Create(fname)
	assert.Nil(t, err)
	return f
}

func Test_NewEndpointsFromFile_WhenFileDoesNotExist(t *testing.T) {
	path, err := ioutil.TempDir("", "tagger-tests")
	assert.Nil(t, err)
	defer os.RemoveAll(path)

	fname := filepath.Join(path, "non_existent.yml")

	endpoints, err := NewEndpointsFromFile(fname)
	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no such file")
}

func Test_NewEndpointsFromFile_WhenUnmarshalFails(t *testing.T) {
	path, err := ioutil.TempDir("", "tagger-tests")
	assert.Nil(t, err)
	defer os.RemoveAll(path)

	fname := filepath.Join(path, "test_file.yml")
	f := TouchFile(t, fname)
	f.Write([]byte("invalid:yml"))
	f.Close()

	endpoints, err := NewEndpointsFromFile(fname)
	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unmarshal errors")
}

func Test_NewEndpointsFromFile_SetupTheEndpoints(t *testing.T) {
	endpoints, err := NewEndpointsFromFile("sample.yml")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(endpoints.Endpoints))
	l := endpoints.Endpoints[0]
	assert.Equal(t, "/hello", l.UriPath)
	assert.Equal(t, "GET", l.Method)
	assert.Equal(t, 200, l.ResponseCode)
	assert.Equal(t, "{msg:\"hello there\"}", l.ResponseBody)
	assert.Equal(t, time.Duration(0), l.Latency)
	assert.Equal(t, 3, len(l.Headers))
	assert.Equal(t, "application/json", l.Headers["Content-Type"])
	assert.Equal(t, "bar", l.Headers["foo"])
	assert.Equal(t, "b", l.Headers["a"])
	l = endpoints.Endpoints[1]
	assert.Equal(t, "/goodbye", l.UriPath)
	assert.Equal(t, "POST", l.Method)
	assert.Equal(t, 201, l.ResponseCode)
	assert.Equal(t, "", l.ResponseBody)
	assert.Equal(t, 100*time.Millisecond, l.Latency)
	assert.Equal(t, 2, len(l.Headers))
	assert.Equal(t, "application/xml", l.Headers["Content-Type"])
	assert.Equal(t, "beta", l.Headers["alpha"])
}

func Test_Endpoints_Validate_When_PathIsEmpty(t *testing.T) {
	endpoints := &Endpoints{Endpoints: []*Endpoint{&Endpoint{}}}
	err := endpoints.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Empty uri_path")
}

func Test_Endpoints_Validate_When_DuplicatePathsExist(t *testing.T) {
	endpoints := &Endpoints{
		Endpoints: []*Endpoint{
			&Endpoint{UriPath: "foo"},
			&Endpoint{UriPath: "foo"},
		},
	}
	err := endpoints.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Path foo already defined")
}
