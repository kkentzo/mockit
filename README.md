[![Travis Build Status](https://travis-ci.org/kkentzo/mockit.svg?branch=master)](https://travis-ci.org/kkentzo/mockit)

# mockit

`mockit` is a command-line utility to quickly mock various aspects of
an http endpoint using a real http server for development
purposes. `mockit` can be used during the development of applications
that consume external http services in order to test the application's
behaviour under various conditions (invalid responses, increased
latencies, various status codes) without the need to actually call
these services.

`mockit` is intentionally dead simple, its main objective being the
quick mocking of an http endpoint using the command line. There exist
[numerous](https://github.com/iridakos/duckrails)
[other](https://github.com/gencebay/httplive)
[programs](https://github.com/jamesdbloom/mockserver) that offer far
more sophisticated features - feel free to check them out if you have
more advanced requirements.

You can start using `mockit` by
[grabbing](https://github.com/kkentzo/mockit/releases) and running the
binary that corresponds to your platform (linux, macOS and windows).

If you have a working Go environment and want to see/play with the
source code, you can also `go get github.com/kkentzo/mockit`.

## Usage

Stuff that can be mocked as command-line arguments to `mockit`
include:

* the URI path (e.g. `-uri /a/random/path`)
* the requested HTTP method (e.g. `-method POST`)
* the server's port (e.g. `-port 9898`)
* the status code of the response (e.g. `-status 401`)
* the payload of the response (e.g. `-body hello`)
* the latency of the response (e.g. `-latency 2s`)
* the response headers (comma-separated, e.g. `-headers
  foo:bar,content-type:application/json`)

Verbose logging is on by default; it can be de-activated using `-verbose=false`.

## Coming up

* support multiple endpoints using `yaml` config file
