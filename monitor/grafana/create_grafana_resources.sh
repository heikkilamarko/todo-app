#!/bin/bash

USER=$1

# datasources

curl \
  -u $USER \
  -d "@datasources/loki.json" \
  -H "Content-Type: application/json" \
  -X POST http://localhost:3000/api/datasources

curl \
  -u $USER \
  -d "@datasources/postgres.json" \
  -H "Content-Type: application/json" \
  -X POST http://localhost:3000/api/datasources

# dashboards

curl \
  -u $USER \
  -d "@dashboards/demo.json" \
  -H "Content-Type: application/json" \
  -X POST http://localhost:3000/api/dashboards/db
