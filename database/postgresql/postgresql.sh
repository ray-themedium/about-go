#!/bin/sh

docker run \
    -d \
    -e POSTGRES_HOST_AUTH_METHOD=trust \
    -e POSTGRES_USER=leo \
    -e POSTGRES_PASSWORD=9036 \
    -e POSTGRES_DB=golang_pg \
    -p 5432:5432 \
    postgres:12.5-alpine