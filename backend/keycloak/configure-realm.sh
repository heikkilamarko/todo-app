#!/bin/bash

path="./configure"
image="todo-app/keycloak-configure"
container="todo-app-keycloak-configure"
env_file="../../env/keycloak-configure.env"
network="todo-app"

docker build -q -t $image $path
docker run --rm --net $network --env-file $env_file --name $container $image
