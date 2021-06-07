#!/bin/bash

# check input arguments

if [[ -z "$1" ]]; then
  echo "error: You must pass a realm json file as the only argument."
  exit 0
fi

# ask keycloak parameters from user

echo "Please type your Keycloak info below"
read -p  "  url: "      KEYCLOAK_URL
read -p  "  username: " KEYCLOAK_USER
read -sp "  password: " KEYCLOAK_PASSWORD
echo -e "\n\nImporting realm...\n\n"


# get access token

TOKEN=$(curl \
  -d "grant_type=password&client_id=admin-cli&username=$KEYCLOAK_USER&password=$KEYCLOAK_PASSWORD" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -X POST $KEYCLOAK_URL/auth/realms/master/protocol/openid-connect/token \
  | jq -r '.access_token')

# import realm

curl \
  -d "@$1" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -X POST $KEYCLOAK_URL/auth/admin/realms
