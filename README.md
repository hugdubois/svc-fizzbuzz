# Simple fizzbuzz microservice

The original fizzbuzz consists in writing all numbers from 1 to 100, and
just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz",
and all multiples of 15 by "fizzbuzz".
The output would look like this:
`"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."`

The goal is to implement a web server that will expose a REST API endpoint
that:

- Accepts five parameters : three integers int1, int2 and limit, and two
strings str1 and str2.

- Returns a list of strings with numbers from 1 to limit, where: all
multiples of int1 are replaced by str1, all multiples of int2 are replaced
by str2, all multiples of int1 and int2 are replaced by str1str2.

## Up and Running

### Baremetal

todo

### Docker

todo

### Kubernetes

todo

## Rest API

todo

## GOAL / ROADMAP

- [x] separate core domain (without dependencies)
- [ ] not too much dependencies
- [x] nice cli with usage and help
- [x] light simple http service (with gracefull shutdown and errors recovering)
- [x] nice requests log
- [x] allows CORS
- [x] endpoint to expose prometheus metrics
- [ ] test coverage >~ 80% (core 100%)
- [ ] good documentation
- [ ] (bonus) add a cache
- [ ] docker / docker-compose
- [ ] simple k8s service
- [ ] (extra bonus) exposes an openapi / swagger file
