#!/bin/bash

# ask keycloak parameters from user

echo "Please type your Keycloak info below"
read -p  "  host: "     KEYCLOAK_HOST
read -p  "  username: " KEYCLOAK_USER
read -sp "  password: " KEYCLOAK_PASSWORD

echo -e "\n\nImporting realm and users...\n\n"

# get access token

TOKEN=$(curl \
  -d "grant_type=password&client_id=admin-cli&username=$KEYCLOAK_USER&password=$KEYCLOAK_PASSWORD" \
  -H "Host: $KEYCLOAK_HOST" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -X POST "$KEYCLOAK_HOST/realms/master/protocol/openid-connect/token" \
  | jq -r '.access_token')

# import realm

curl \
  -d "@realms/todo-app.json" \
  -H "Host: $KEYCLOAK_HOST" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -X POST "$KEYCLOAK_HOST/admin/realms"

# import users

curl \
  -d "@users/demouser.json" \
  -H "Host: $KEYCLOAK_HOST" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -X POST "$KEYCLOAK_HOST/admin/realms/todo-app/users"

curl \
  -d "@users/demoviewer.json" \
  -H "Host: $KEYCLOAK_HOST" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -X POST "$KEYCLOAK_HOST/admin/realms/todo-app/users"
