#!/usr/bin/env bash

export $(grep -v '^#' ../../../todo-app-secrets/postgres_migrate.env | xargs)

docker run \
  --rm \
  --mount type=bind,src=$(pwd)/migrations,dst=/migrations \
  migrate/migrate -path /migrations \
  -database $POSTGRES_CONNECTIONSTRING \
  up
