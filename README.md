# Todo App

![Todo App Architecture](doc/architecture.png)

## Running the app

```bash
# Build and run the app
> docker compose up --build
```

App URL: http://localhost:8000

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
