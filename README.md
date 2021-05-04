# Todo App

![Todo App Architecture](doc/architecture.png)

## Running the app

```bash
# Tested with Docker Desktop (Mac and Windows)

# Build all docker images and run the app
> docker compose up --build

# Run database migrations
> cd backend/db
> docker compose run --rm migrate

# Open browser and navigate to the app url
> open http://localhost:8000
```
