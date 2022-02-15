#!/bin/bash

# check input arguments

if [[ -z "$1" ]]; then
  echo "error: You must pass a realm json file as the only argument."
  exit 0
fi

# ask keycloak parameters from user

echo "Please type your Keycloak info below"
read -p  "  realm: "    KEYCLOAK_REALM
read -p  "  url: "      KEYCLOAK_URL
read -p  "  username: " KEYCLOAK_USER
read -sp "  password: " KEYCLOAK_PASSWORD

echo -e "\n\nExporting realm...\n\n"

# get access token

TOKEN=$(curl \
  -d "grant_type=password&client_id=admin-cli&username=$KEYCLOAK_USER&password=$KEYCLOAK_PASSWORD" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -X POST "$KEYCLOAK_URL/realms/master/protocol/openid-connect/token" \
  | jq -r '.access_token')

# export realm

curl -o $1 \
  -H "Authorization: Bearer $TOKEN" \
  -X POST "$KEYCLOAK_URL/admin/realms/$KEYCLOAK_REALM/partial-export?exportClients=true&exportGroupsAndRoles=true"
