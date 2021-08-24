#!/bin/sh

USER=$GF_USER:$GF_PASSWORD

# datasources

curl \
  -u $USER \
  -d "$(envsubst < datasources/loki.json)" \
  -H "Content-Type: application/json" \
  -X POST $GF_URL/api/datasources

curl \
  -u $USER \
  -d "$(envsubst < datasources/postgres.json)" \
  -H "Content-Type: application/json" \
  -X POST $GF_URL/api/datasources

# dashboards

curl \
  -u $USER \
  -d "$(envsubst < dashboards/demo.json)" \
  -H "Content-Type: application/json" \
  -X POST $GF_URL/api/dashboards/db
