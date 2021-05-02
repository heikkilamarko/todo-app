#!/usr/bin/env bash

export $(grep -v '^#' ../../env/migrate.env | xargs)

docker run \
  --rm \
  --mount type=bind,src=$(pwd)/migrations,dst=/migrations \
  migrate/migrate -path /migrations \
  -database $POSTGRES_CONNECTIONSTRING \
  up
