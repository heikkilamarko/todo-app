#!/bin/bash

command="$1"

command_regex="^(import|export)$"

if [[ ! "$command" =~ $command_regex ]]; then
	cat <<- EOF
		usage:
		    $0 <command>

		    Imports and exports keycloak realms.

		command:
		    import
		    export
	EOF
	exit 1
fi

path="./configure"
image="todo-app/keycloak-configure"
container="keycloak-configure"
env_file="../../env/keycloak.env"
network="todo-app"
realm="todo-app"

function import_realm() {
	docker run --rm --net $network --env-file $env_file --name $container $image "import"
}

function export_realm() {
	docker run --net $network --env-file $env_file --name $container $image "export"
	docker cp $container:/$realm.json ./$realm.json
	docker rm $container
}

docker build -q -t $image $path

${command}_realm

docker image rm $image
