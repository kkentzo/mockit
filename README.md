[![Travis Build Status](https://travis-ci.org/kkentzo/mockit.svg?branch=master)](https://travis-ci.org/kkentzo/mockit)

# mockit

A utility to quickly mock various aspects of an http endpoint for
development purposes using a real http server.

## Details

Stuff that can be mocked as command-line arguments to `mockit`
include:

* the URI path (e.g. `-uri /a/random/path`)
* the requested HTTP method (e.g. `-method POST`)
* the server's port (e.g. `-port 9898`)
* the status code of the response (e.g. `-status 401`)
* the latency of the response (e.g. `-latency 2s`)
* the response headers (comma-separated, e.g. `-headers
  foo:bar,content-type:application/json`)

Verbose logging is on by default; it can be de-activated using `-verbose=false`.

## Coming up

* custom response payloads
* support multiple endpoints using `yaml` config file

## Installation

Ready-made binaries
are [available](https://github.com/kkentzo/mockit/releases)
(linux-only for the moment - other platforms coming up).

If you have a working Go environment, you can also `go get
github.com/kkentzo/mockit`.
