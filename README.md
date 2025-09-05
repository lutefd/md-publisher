# Self-Hosted Obsidian Publish Alternative

A self-hosted alternative to Obsidian Publish that allows you to publish your Markdown notes to your own website.

## Architecture

### Components

1. **Go API**

   - Endpoints for publishing and unpublishing notes
   - BadgerDB for key-value storage
   - Markdown export functionality
   - Queue system for debouncing rebuilds

2. **SvelteKit Frontend**

   - Server-side rendering and static site generation capabilities
   - Svelte components for interactive features
   - Tailwind CSS for styling
   - Client-side search functionality

3. **Web Server**
   - Caddy configuration for serving the application

## Project Structure

```
/publisher
  /api
    /cmd
      /server        # Main server entry point
    /internal
      /api           # API handlers
      /storage       # BadgerDB integration
    /data            # BadgerDB files
  /web
    /src
      /lib           # Shared libraries and components
      /routes        # SvelteKit routes
      /app.css       # Global styles
    /static          # Static assets
  Caddyfile.sample   # Sample Caddy web server configuration
```

## Getting Started

### Prerequisites

- Go 1.25+
- Node.js 22+
- Caddy (for production deployment)

### Development Setup

1. **Clone the repository**

```bash
git clone https://github.com/lutefd/publisher.git
cd publisher
```

2. **Set up the Go API**

```bash
cd api
go mod tidy
go run cmd/server/main.go
```

3. **Set up the SvelteKit frontend**

```bash
cd ../web
npm install
npm run dev
```

4. **Build for production**

```bash
cd ../web
npm run build
```

5. **Configure Caddy**

Copy the `Caddyfile.sample` to `Caddyfile` and update it with your domain and the correct paths.

```bash
cp Caddyfile.sample Caddyfile
# Edit the Caddyfile with your settings
caddy run
```

## Usage

### Publishing Notes

To publish a note, send a POST request to the API:

```bash
curl -X POST http://localhost:8080/publish \
  -H "Content-Type: application/json" \
  -d '{
    "id": "my-note",
    "content": "# My Note\n\nThis is my note content.",
    "metadata": {
      "title": "My Note",
      "tags": ["example", "note"]
    }
  }'
```

### Unpublishing Notes

To unpublish a note, send a DELETE request to the API:

```bash
curl -X DELETE http://localhost:8080/note/my-note
```

## License

MIT
