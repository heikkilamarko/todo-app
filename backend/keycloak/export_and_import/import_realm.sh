#!/bin/bash

KEYCLOAK_URL=http://localhost:8002
KEYCLOAK_USER=admin
KEYCLOAK_PASSWORD=admin

# Get access token

TOKEN=$(curl \
  -d "grant_type=password&client_id=admin-cli&username=$KEYCLOAK_USER&password=$KEYCLOAK_PASSWORD" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -X POST $KEYCLOAK_URL/auth/realms/master/protocol/openid-connect/token \
  | jq -r '.access_token')

# Import realm

curl \
  -i \
  -d "@$1" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -X POST $KEYCLOAK_URL/auth/admin/realms
