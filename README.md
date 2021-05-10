# Todo App

![Todo App Architecture](doc/architecture.png)

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
# Build
> docker compose build

# Run NATS container
> docker compose up nats

# Configure NATS JetStream
> cd <repo>/backend/nats/jetstream
> ./configure_jetstream.sh

# Run all containers
> docker compose up

# Configure Grafana Dashboard
#   <username>: grafana admin username (default: admin)
#   <password>: grafana admin password (default: admin)
> cd <repo>/monitor/grafana
> ./create_grafana_resources.sh <username>:<password>
```

App: http://localhost:8000

Grafana: http://localhost:3000

MinIO: http://localhost:9000
