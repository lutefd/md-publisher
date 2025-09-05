# Docker Setup for Publisher

This document explains how to run the Publisher project using Docker.

## Prerequisites

- Docker and Docker Compose installed on your system
- Caddy server running with a network named `caddy_net`

## Components

The project consists of two main services:

1. **Frontend (SvelteKit)**: A modern web framework that serves the published notes
2. **API (Go)**: A backend service that handles the publishing workflow and uses BadgerDB for storage

## Configuration

### Docker Compose

The `docker-compose.yml` file defines the services and their configurations:

```yaml
version: "3.8"

services:
  publisher-frontend:
    build:
      context: ./web
      dockerfile: Dockerfile
    ports:
      - "5173:5173"
    environment:
      - NODE_ENV=production
    depends_on:
      - pubslisher-api
    networks:
      - caddy_net
    restart: unless-stopped

  pubslisher-api:
    build:
      context: ./api
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    volumes:
      - badger_data:/app/data
    environment:
      - GIN_MODE=release
    networks:
      - caddy_net
    restart: unless-stopped

networks:
  caddy_net:
    external: true

volumes:
  badger_data:
    driver: local
```

### Security

Both services use distroless container images to minimize attack surface:

- The Go API uses `gcr.io/distroless/static-debian12`
- The SvelteKit frontend uses Node.js with a minimal container footprint

### Data Persistence

- BadgerDB data is stored in a named Docker volume (`badger_data`)

## Usage

### Starting the Services

```bash
docker compose up -d
```

### Stopping the Services

```bash
docker compose down
```

### Viewing Logs

```bash
# All services
docker compose logs

# Specific service
docker compose logs frontend
docker compose logs api
```

### Rebuilding Services

```bash
docker compose build
```

## Integration with Caddy

The services connect to an external Caddy network (`caddy_net`). Make sure your Caddy configuration includes the appropriate routes to these services.

Example Caddyfile snippet:

```
publisher.example.com {
    reverse_proxy publisher-frontend:5173
}

api.publisher.example.com {
    reverse_proxy pubslisher-api:8080
}
```

## Volumes

- `badger_data`: Persistent storage for BadgerDB
