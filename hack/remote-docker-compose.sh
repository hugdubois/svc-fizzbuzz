#!/bin/sh

BRANCH=develop
BASE_URL=https://raw.githubusercontent.com/hugdubois/svc-fizzbuzz/$BRANCH
SVC_TAG=1.0.0

mkdir -p svc-fizzbuzz-compose
cd svc-fizzbuzz-compose
mkdir -p infra
cd infra
curl -O $BASE_URL/infra/config.monitoring
mkdir -p prometheus
cd prometheus
curl -O $BASE_URL/infra/prometheus/prometheus.yml
cd ../..
curl -O $BASE_URL/docker-compose.yml

echo "TAG=$SVC_TAG" > .env
docker-compose -d up
