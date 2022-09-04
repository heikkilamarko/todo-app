#!/bin/sh

# get access token

token=$(curl \
  -d "grant_type=password&client_id=admin-cli&username=$KEYCLOAK_ADMIN&password=$KEYCLOAK_ADMIN_PASSWORD" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -X POST "$KEYCLOAK_URL/realms/master/protocol/openid-connect/token" \
  | jq -r '.access_token')

# import realms

for f in $(find realms -maxdepth 2 -type f); do
  curl \
    -d "@$f" \
    -H "Authorization: Bearer $token" \
    -H "Content-Type: application/json" \
    -X POST "$KEYCLOAK_URL/admin/realms"
done;

# import users

for f in $(find users -maxdepth 2 -type f); do
  curl \
    -d "@$f" \
    -H "Authorization: Bearer $token" \
    -H "Content-Type: application/json" \
    -X POST "$KEYCLOAK_URL/admin/realms/$KEYCLOAK_REALM/users"
done;
