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
# Build and run
> docker compose up --build
```

App: http://localhost:8000

Grafana: http://localhost:3000

MinIO: http://localhost:9000
