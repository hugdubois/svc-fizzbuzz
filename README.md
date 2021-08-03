# Simple fizzbuzz microservice

[![Build Status](https://travis-ci.com/hugdubois/svc-fizzbuzz.svg?branch=develop)](https://travis-ci.com/hugdubois/svc-fizzbuzz)
[![codecov](https://codecov.io/gh/hugdubois/svc-fizzbuzz/branch/develop/graph/badge.svg?token=E6E9CSRY80)](https://codecov.io/gh/hugdubois/svc-fizzbuzz)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hugdubois/svc-fizzbuzz)](https://pkg.go.dev/github.com/hugdubois/svc-fizzbuzz)
[![License: MIT](https://img.shields.io/badge/License-MIT-violet.svg)](https://opensource.org/licenses/MIT)

## Table of contents

- [General info](#general-info)
- [Technologies](#technologies)
- [Up and running](#up-and-running)
- [Usage](#usage)
- [Examples](#examples)
- [Contribute](#contribute)
- [Roadmap](#roadmap)
- [Authors](#authors)
- [Support](#support)
- [License](#license)

## General info

The original fizzbuzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz".  The output would look like this:

```
"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."
```

The goal is to implement a web server that will expose a REST API endpoint that:

- Accepts five parameters : three integers `int1`, `int2` and `limit`, and two strings `str1` and `str2`.
- Returns a list of strings with numbers from 1 to `limit`, where: all multiples of `int1` are replaced by `str1`, all multiples of `int2` are replaced by `str2`, all multiples of `int1` and `int2` are replaced by `str1str2`.

The server needs to be:

- Ready for production
- Easy to maintain by other developers
- Add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
    - Accept no parameter
    - Return the parameters corresponding to the most used request, as well as the number of hits for this request

The __svc-fizzbuzz__ microservice exposes a REST API with the following endpoints:

- __/__
   - Forwards to __/version__ endpoint.

- __/version__
   - Returns the service version.

- __/status__
   - Returns the service status.

- __/metrics__
   - Returns the __prometheus__ metrics.

- __/api/v1/fizzbuzz__
   - Returns a list of strings with numbers from 1 to `limit`, where: all multiples of `int1` are replaced by `str1`, all multiples of `int2` are replaced by `str2`, all multiples of `int1` and `int2` are replaced by `str1str2`.
   - Query String (or __POST__ body parameters):
       - `linit` (__positive integer__) max value `100 000`
       - `int1` (__positive integer__) default value `3`
       - `int2` (__positive integer__) default value `5`
       - `str1` (__string__) default value `fizz`
       - `str2` (__string__) default value `buzz`

- __/api/v1/fizzbuzz/top__
   - Returns usage statistics of the __/api/v1/fizzbuzz__ endpoint. It allows the users to know what the number of hits of that endpoint. And returns the parameters corresponding to it.

## Technologies

Project is created with:

- __go__: 1.16
- __redis__: 6.0.9
- __docker__: 20.10.6

## Up and Running

⚠️ To __baremetal__ and __docker__ methods below, a redis-server is needed.
To set the connection use `--database-connect` flag see [Usage](#usage) for more information.

### Baremetal

Install localy (__baremetal__) (needs a redis server) :

    $ go get github.com/hugdubois/svc-fizzbuzz

### Docker

Install via __Docker__ (needs a redis server) :

    $ docker pull hugdubois/svc-fizzbuzz:0.0.6_dev
    $ docker run -d --name=svc-fizzbuzz --net=host -it hugdubois/svc-fizzbuzz:0.0.6_dev serve --database-connect localhost:6379

__NOTE__ : __docker__ images can be found on [dockerhub](https://hub.docker.com/r/hugdubois/svc-fizzbuzz).

### Docker compose

Install via __docker-compose__ (without git clone) (needs curl) :

    $ curl https://raw.githubusercontent.com/hugdubois/svc-fizzbuzz/master/hack/remote-docker-compose.sh | sh

__NOTE__ : that script will create a __svc-fizzbuzz-compose__ with all required files

With `git clone` simply do :

    $ git clone https://github.com/hugdubois/svc-fizzbuzz
    $ cd svc-fizzbuzz
    $ make compose-up

Some services are exposed:

- __svc-fizzbuzz__ on http://localhost:8080/
- __prometheus__ on http://localhost:9090/
- __grafana__ on http://localhost:3000/
  - some dashboards can be found [here](./infra/grafana/dashboards)

### Kubernetes

Install via __kubernetes__ (needs kubectl):

    $ kubectl apply -f https://raw.githubusercontent.com/hugdubois/svc-fizzbuzz/master/k8s-deployment.yaml

__NOTE__: if you use __minikube__ do `minikube service svc-fizzbuzz` to expose and get the service ip.

## Usage

__svc-fizzbuzz__ is a sipmle fizzbuzz microservice.

Basic usage:

    $ svc-fizzbuzz [command]

Available Commands:

- __completion__ generates the autocompletion script for the specified shell
- __help__ help about any command
- __serve__ launches the svc-fizzbuzz service webserver
- __version__ returns service version

To get __help__ simply run `svc-fizzbuzz help`.

To launch the __API webserver__ run: `svc-fizzbuzz serve`

- Some flags are available :
   - __--address__ (string) (short __-a__)
      - Must be used to set the HTTP server address.
      - ex: `127.0.0.1:13000`
      - default: `:8080`
   - __--cors-origin__ (string) (short __-c__)
      - Must be used to set the _Cross Origin Resource Sharing AllowedOrigins_. It's a string separed by `|`.
      - ex: `http://*domain1.com|http://*domain2.com`
      - default: `*`
   - __--database-connect__ (string)
      - Must be used to set the redis server connection informations. __[[db:]password@]host:port__.
      - ex: `1:passW0rd@redis-server:6379`
      - default: `localhost:6379`
   - __--debug__ (boolean)
      - Must be used to force debug mode.
      - ex: `1:passW0rd@redis-server:6379`
      - default: `localhost:6379`
   - __--help__ (string)
      - Must be used to get help.
   - __--read-timeout__ (duration)
      - Must be used to set the server read timeout (5s,5m,5h) before connection is cancelled.
      - ex: `10s`
      - default: `5s`
   - __--shutdown-timeout__ (duration)
      - Must be used to set the server shutdown timeout (5s,5m,5h) graceful shutdown.
      - ex: `15s`
      - default: `10s`
   - __--write-timeout__ (duration)
      - Must be used to set the server write timeout (5s,5m,5h) before connection is cancelled.
      - ex: `15s`
      - default: `10s`

- The __svc-fizzbuzz__ microservice exposes a REST API with the following endpoints:
   - __/__
      - Forwards to __/version__ endpoint.
   - __/version__
      - Returns the service version.
   - __/status__
      - Returns the service status.
   - __/metrics__
      - Returns the __prometheus__ metrics.
   - __/api/v1/fizzbuzz__
      - Returns a list of strings with numbers from 1 to `limit`, where: all multiples of `int1` are replaced by `str1`, all multiples of `int2` are replaced by `str2`, all multiples of `int1` and `int2` are replaced by `str1str2`.
      - Query String (or __POST__ body parameters):
          - `linit` (__positive integer__) max value `100 000`
          - `int1` (__positive integer__) default value `3`
          - `int2` (__positive integer__) default value `5`
          - `str1` (__string__) default value `fizz`
          - `str2` (__string__) default value `buzz`
   - __/api/v1/fizzbuzz/top__
      - Returns usage statistics of the __/api/v1/fizzbuzz__ endpoint. It allows the users to know what the number of hits of that endpoint. And returns the parameters corresponding to it.

## Examples

### /

```shell
$ curl "localhost:8080/"
```

Should return

```json
{"name":"svc-fizzbuzz","version":"v0.0.5"}
```

### /status

```shell
$ curl "localhost:8080/status"
```

Should return

```json
{"svc-alive":true,"store-alive":true}
```

### /api/v1/fizzbuzz

This is the core API endpoint. This endpoint returns a list of strings with numbers from 1 to `limit`, where: all multiples of `int1` are replaced by `str1`, all multiples of `int2` are replaced by `str2`, all multiples of `int1` and `int2` are replaced by `str1str2`.

```shell
$ curl "localhost:8080/api/v1/fizzbuzz"
```

Should return a original fizzbuzz

```json
{"fizzbuzz":["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31","32","fizz","34","buzz","fizz","37","38","fizz","buzz","41","fizz","43","44","fizzbuzz","46","47","fizz","49","buzz","fizz","52","53","fizz","buzz","56","fizz","58","59","fizzbuzz","61","62","fizz","64","buzz","fizz","67","68","fizz","buzz","71","fizz","73","74","fizzbuzz","76","77","fizz","79","buzz","fizz","82","83","fizz","buzz","86","fizz","88","89","fizzbuzz","91","92","fizz","94","buzz","fizz","97","98","fizz","buzz"]}
```

The query string (or __POST__ body parameters):

  - `linit` (__positive integer__) max value `100 000`
  - `int1` (__positive integer__) default value `3`
  - `int2` (__positive integer__) default value `5`
  - `str1` (__string__) default value `fizz`
  - `str2` (__string__) default value `buzz`

⚠️ All of them are optional, if missing the default value is considered.

```shell
$ curl "localhost:8080/api/v1/fizzbuzz?limit=10"
```

Should return only ten values of the original fizzbuzz.

```json
{"fizzbuzz":["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz"]}
```

More complex query:

```shell
$ curl "localhost:8080/api/v1/fizzbuzz?limit=10&int1=2&int3=4&str1=bon&str2=coin"
```

Should return only ten values of a custom fizzbuzz.

```json
{"fizzbuzz":["1","bon","3","bon","coin","bon","7","bon","9","boncoin"]}
```

⚠️ There is no restriction on the HTTP verb so __POST__, __PUT__, __PATCH__ methods are accepted and valid. It is the same to all endpoints.

So all of these calls

```shell
$ curl -XPOST "localhost:8080/api/v1/fizzbuzz?limit=10&int1=2&int3=4&str1=bon&str2=coin"
$ curl -XPUT "localhost:8080/api/v1/fizzbuzz?limit=10&int1=2&int3=4&str1=bon&str2=coin"
$ curl -XPATCH "localhost:8080/api/v1/fizzbuzz?limit=10&int1=2&int3=4&str1=bon&str2=coin"
$ curl -XPOST -d "limit=10&int1=2&int3=4&str1=bon&str2=coin" "localhost:8080/api/v1/fizzbuzz"
$ curl -XPUT -d "limit=10&int1=2&int3=4&str1=bon&str2=coin" "localhost:8080/api/v1/fizzbuzz"
$ curl -XPATCH -d "limit=10&int1=2&int3=4&str1=bon&str2=coin" "localhost:8080/api/v1/fizzbuzz"
```

Should return only ten values of a custom fizzbuzz.

```json
{"fizzbuzz":["1","bon","3","bon","coin","bon","7","bon","9","boncoin"]}
```

⚠️ If bad parameters are send, an error is returned with the `422 Unprocessable Entity` with the reason of the error.

So :

```shell
$ curl -v "localhost:8080/api/v1/fizzbuzz?limit=infini"
```

Should return an error.

```json
...
< HTTP/1.1 422 Unprocessable Entity
...
{"code":422,"message":"Bad parameter: 'limit' must be a positive number - got (infini)"}%
```

And

```shell
$ curl "localhost:8080/api/v1/fizzbuzz?int1=bon"
```

Should return also an error.

```json
{"code":422,"message":"Bad parameter: 'int1' must be a positive number - got (bon)"}%
```

⚠️ Only one error is returned

So :

```shell
curl "localhost:8080/api/v1/fizzbuzz?int1=bad&int2=bad"
```

Should return only one error.

```json
{"code":422,"message":"Bad parameter: 'int1' must be a positive number - got (bad)"}%
```


### /api/v1/fizzbuzz/top

This endpoint returns the usage statistics of the __/api/v1/fizzbuzz__ endpoint. It allows the users to know what the number of hits of that endpoint. And returns the parameters corresponding to it.

```shell
$ curl "localhost:8080/api/v1/fizzbuzz/top"
```

Should return only one error.

```json
{"data":{"params":{"limit":10,"int1":2,"str1":"bon","int2":5,"str2":"coin"},"count_request":10}}%
```

### /metrics

This endpoint exposes the prometheus metrics.

```shell
$ curl "localhost:8080/metrics"
```

Should return.

```json
...
# HELP fizzbuzz_http_requests_duration_millisecond How long it took to process the request, partitioned by status code, method and HTTP path.
# TYPE fizzbuzz_http_requests_duration_millisecond summary
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="GET",path="/"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="GET",path="/"} 1
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="GET",path="/api/v1/fizzbuzz"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="GET",path="/api/v1/fizzbuzz"} 3
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="GET",path="/api/v1/fizzbuzz/top"} 1
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="GET",path="/api/v1/fizzbuzz/top"} 1
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="GET",path="/status"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="GET",path="/status"} 1
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="HEAD",path="/api/v1/fizzbuzz"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="HEAD",path="/api/v1/fizzbuzz"} 1
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="OPTIONS",path="/api/v1/fizzbuzz"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="OPTIONS",path="/api/v1/fizzbuzz"} 1
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="PATCH",path="/api/v1/fizzbuzz"} 4
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="PATCH",path="/api/v1/fizzbuzz"} 2
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="POST",path="/api/v1/fizzbuzz"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="POST",path="/api/v1/fizzbuzz"} 2
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="POST",path="/api/v1/fizzbuzz/top"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="POST",path="/api/v1/fizzbuzz/top"} 1
fizzbuzz_http_requests_duration_millisecond_sum{code="OK",method="PUT",path="/api/v1/fizzbuzz"} 4
fizzbuzz_http_requests_duration_millisecond_count{code="OK",method="PUT",path="/api/v1/fizzbuzz"} 5
fizzbuzz_http_requests_duration_millisecond_sum{code="Unprocessable Entity",method="GET",path="/api/v1/fizzbuzz"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="Unprocessable Entity",method="GET",path="/api/v1/fizzbuzz"} 6
fizzbuzz_http_requests_duration_millisecond_sum{code="Unprocessable Entity",method="PATCH",path="/api/v1/fizzbuzz"} 0
fizzbuzz_http_requests_duration_millisecond_count{code="Unprocessable Entity",method="PATCH",path="/api/v1/fizzbuzz"} 1
...
```

### Not found

All others calls return an `404 Not Found` error.

```shell
$ curl -v "localhost:8080/does_not_exists"
```

Should return.

```json
...
< HTTP/1.1 404 Not Found
...
{"code":404,"message":"Not Found"}
```

## Contribute

This repository follows the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) and standard [git flow](https://danielkummer.github.io/git-flow-cheatsheet/).

See package documentation on [pkg.go.dev](https://pkg.go.dev/github.com/hugdubois/svc-fizzbuzz).

Code coverage can be found on [codecov](https://app.codecov.io/gh/hugdubois/svc-fizzbuzz).

CI is on [travis](https://travis-ci.com/github/hugdubois/svc-fizzbuzz) and [github actions](https://github.com/hugdubois/svc-fizzbuzz/actions).

### Make directives

- `make build` - (default) build the service and inject the version (`-ldflags`).
- `make version` - Display current version (__VERSION__ file).
- `make test` - Run test.
- `make test-v` - Run test on verbose mode.
- `make test-live` - Run test on infinite shell loop.
- `make test-cover` - Run test with coverage.
- `make test-cover-profile` - Run test with coverage and generate a profile coverage file.
- `make test-cover-report` - Run test with coverage UI on a browser (`go tool cover -html=...`).
- `make test-cover-func` - Run test with total coverage computation (`go tool cover -func=...`).
- `make serve` - Build and launch the server api with the debug mode activate.
- `make clean` - Removing all generated files (compiled files, code coverage, ...).
- `make docker-tag` - Generate `.env` file to __docker-compose__.
- `make docker` - Generate __docker__ image.
- `make docker-push` - Push the __docker__ image to the repository (use `DOCKER_REGISTRY` and `DOCKER_IMAGE_NAME` like this `DOCKER_REGISTRY={{hostname}}:{{port}} make docker-push`)
- `make docker-run` - Run service with __docker__.
- `make docker-rm` - Remove service from __docker__.
- `make compose-up` - Run `docker-compose up -d`.
- `make compose-down` - Run `docker-compose up`.
- `make compose-ps` -  Run `docker-compose ps`.
- `make k8s-deploy` - Deploy __kubernetes__ deployment on the current cluster via `kubectl`.
- `make k8s-delete` -  Delete __kubernetes__ deployment from the current cluster via `kubectl`.
- `make update-pkg-cache` - Performs a `go get` via https://proxy.golang.org  as `GOPROXY`.

### Directories structure

- `cmd` - The directory contains the `cli` commands.
- `core` - The directory is the core domain layer.
- `hack` - The directory contains some shell scripts.
- `helpers` - The directory contains a helpers package.
- `infra` - The directory contains some infrastructure code to __docker-compose__.
- `middlewares` - The directory contains all HTTP middlewares.
- `service` - The directory is the service layer.
- `store` - The directory is persistence layer.
- `vendor` - The directory is `go mod` vendoring.

### Notes

The statistics to __/api/v1/fizzbuzz/top__ endpoint are stored in a [redis sorted sets](https://redislabs.com/ebook/part-1-getting-started/chapter-1-getting-to-know-redis/1-2-what-redis-data-structures-look-like/1-2-5-sorted-sets-in-redis/).

To retrieve the most used request a simple [ZREVRANGE k 0 0 WITHSCORES](https://redis.io/commands/zrevrange) is good enough.

## Roadmap

- [x] separate core domain (without dependencies)
- [x] nice cli with usage and help
- [x] light simple http service (with graceful shutdown and errors recovering)
- [x] nice requests log
- [x] allows CORS
- [x] endpoint to expose prometheus metrics
- [x] CI
- [x] test coverage >~ 80% (core 100%)
- [x] docker / docker-compose
- [x] simple k8s deployment
    - [ ] use kustomize to bump TAG in k8s-deployment.yaml
- [ ] add openapi
- [ ] add a cache
- [ ] TLS support

## Authors

[Hugues Dubois](https://www.linkedin.com/in/huguesdubois)

## Support

mail: [hugdubois@gmail.com](mailto:hugdubois@gmail.com)

## License

[MIT](./LICENSE)
