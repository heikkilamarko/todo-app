#!/bin/sh

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

token=$(curl \
	-d "grant_type=password&client_id=admin-cli&username=$KEYCLOAK_ADMIN&password=$KEYCLOAK_ADMIN_PASSWORD" \
	-H "Content-Type: application/x-www-form-urlencoded" \
	-X POST "$KEYCLOAK_URL/realms/master/protocol/openid-connect/token" \
	| jq -r '.access_token')


function import_realm() {
	# realm
	curl \
		-d "@realms/$KEYCLOAK_REALM.json" \
		-H "Authorization: Bearer $token" \
		-H "Content-Type: application/json" \
		-X POST "$KEYCLOAK_URL/admin/realms"

	# users
	for f in $(find users -maxdepth 2 -type f); do
		curl \
			-d "@$f" \
			-H "Authorization: Bearer $token" \
			-H "Content-Type: application/json" \
			-X POST "$KEYCLOAK_URL/admin/realms/$KEYCLOAK_REALM/users"
	done;
}

function export_realm() {
	curl \
		-o "/$KEYCLOAK_REALM.json" \
		-H "Authorization: Bearer $token" \
		-X POST "$KEYCLOAK_URL/admin/realms/$KEYCLOAK_REALM/partial-export?exportClients=true&exportGroupsAndRoles=true"
}

${command}_realm
