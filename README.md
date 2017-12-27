# mockit

A utility to quickly mock various aspects of an http endpoint for
development purposes using a real http server.

## Details

Stuff that can be mocked as command-line arguments to `mockit`
include:

* the URI path (e.g. `-path /a/random/path`)
* the server's port (e.g. `-port 9898`)
* the status code of the response (e.g. `-status 401`)
* the latency of the response (e.g. `-latency 2s`)
* the response headers (e.g. `-header foo:bar,'content-type:application/json'`)

## Coming up

* http redirect (temporary, permanent)
* custom response payloads
* more verbose output

## Installation

If you have a working Go environment, use `go get
github.com/kkentzo/mockit`

Ready-made binaries for download coming up...
