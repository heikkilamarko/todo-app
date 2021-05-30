# Todo App

## Application Architecture

![Todo App Architecture](doc/architecture.png)

## Identity and Access Management

- [Keycloak](https://www.keycloak.org/)
  - URL: http://localhost:8002
  - username: `admin`
  - password: `admin`

## Monitoring Stack

- [Grafana](https://grafana.com/oss/grafana/)
  - URL: http://localhost:3000
  - username: `admin`
  - password: `admin`
- [Loki](https://grafana.com/oss/loki/)
- [Promtail](https://grafana.com/docs/loki/latest/clients/promtail/)
- [MinIO](https://min.io/)
  - URL: http://localhost:9000
  - username: `minio`
  - password: `minio123`

## Running the App

```bash
> docker compose up --build
```

App: http://localhost:8000

- username: `demouser`
- password: `demouser`

## Generating JSON Schemas

The app uses JSON Schemas for message validation. Schemas are generated from [AsyncAPI](https://www.asyncapi.com/) documents.

```bash
> cd tools/json-schema-generator
> npm i
> ./generate.sh
```
