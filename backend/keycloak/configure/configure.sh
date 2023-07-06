#!/bin/sh
set -e

# get access token

token=$(curl -X POST \
	-H "Content-Type: application/x-www-form-urlencoded" \
	-d "grant_type=password&client_id=admin-cli&username=$KEYCLOAK_ADMIN&password=$KEYCLOAK_ADMIN_PASSWORD" \
	"$KEYCLOAK_URL/realms/master/protocol/openid-connect/token" \
	| jq -r '.access_token')

# import realm

curl -X POST \
	-H "Authorization: Bearer $token" \
	-H "Content-Type: application/json" \
	-d "@realms/$KEYCLOAK_REALM.json" \
	"$KEYCLOAK_URL/admin/realms"

# import users

if [ -d "users" ]; then
	for f in $(find users -maxdepth 2 -type f); do
		curl -X POST \
			-H "Authorization: Bearer $token" \
			-H "Content-Type: application/json" \
			-d "@$f" \
			"$KEYCLOAK_URL/admin/realms/$KEYCLOAK_REALM/users"
	done;
fi
