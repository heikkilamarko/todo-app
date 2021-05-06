# Todo App

![Todo App Architecture](doc/architecture.png)

## Running the app

```bash
# Tested with Docker Desktop (Mac and Windows)

# Create database and run migrations
> docker compose -f docker-compose.yml -f docker-compose.migrate.yml run --rm migrate

# Build and run the app
> docker compose up --build

# Open the app in browser
> open http://localhost:8000
```
