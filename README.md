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

`mockit` works by defining one or more listeners each characterized by
the certain properties. A single listener can be specified using the
command line using the following arguments:

* URI path (e.g. `-uri /a/random/path`)
* requested HTTP method (e.g. `-method POST`)
* status code of the response (e.g. `-status 401`)
* payload of the response (e.g. `-body hello`)
* latency of the response (e.g. `-latency 2s`)
* response headers (comma-separated, e.g. `-headers
  foo:bar,content-type:application/json`)

The port to which the server binds can be specified using `-port`
(e.g. `-port 8888`). The default value is 9999.

Multiple listeners on the same port are supported through a `yml`
configuration file (see example [here](sample.yml)). The configuration
file path can be passed to `mockit` using the `-config PATH_TO_FILE`
argument.
