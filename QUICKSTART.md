# Quick Start Guide

Get the monorepo up and running in minutes!

## Prerequisites

- Node.js 18+
- Python 3.8+
- Go 1.21+
- Make (optional)

## Installation

### Option 1: Using Make (Recommended)

```bash
make setup
```

This installs all dependencies for all three services.

### Option 2: Manual Setup

```bash
# API (TypeScript/Node.js)
cd api
npm install
cd ..

# Worker (Python)
cd worker
python3 -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate
pip install -r requirements.txt
cd ..

# CLI (Go)
cd cli
go mod download
cd ..
```

## Running Services

Open three terminal windows:

### Terminal 1: API Server

```bash
cd api
npm run dev
```

Server will start on `http://localhost:3000`

### Terminal 2: Worker Service

```bash
cd worker
source venv/bin/activate  # Windows: venv\Scripts\activate
python worker.py
```

### Terminal 3: CLI Tool

```bash
cd cli
go run main.go status
```

## Test It Out

### Create a Task (via curl)

```bash
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn the monorepo"}'
```

### List Tasks

```bash
curl http://localhost:3000/tasks
```

### Check Service Status

```bash
cd cli
go run main.go status
```

### View Monorepo Info

```bash
cd cli
go run main.go info
```

### List All Apps

```bash
cd cli
go run main.go list
```

## Building

To build all services:

```bash
make build
```

Or individually:

```bash
# API
cd api && npm run build

# CLI
cd cli && go build -o monorepo main.go
```

## Docker Setup

To run everything with Docker:

```bash
docker-compose up
```

This will start all three services in containers.

## Project Structure

```
monorepo/
├── api/              # Node.js REST API
│   ├── src/
│   ├── package.json
│   └── tsconfig.json
├── worker/           # Python job processor
│   ├── worker.py
│   ├── requirements.txt
│   └── Dockerfile
├── cli/              # Go CLI tool
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── shared/           # Shared docs
├── Makefile          # Build commands
└── docker-compose.yml
```

## Useful Commands

```bash
# Install everything
make setup

# Build all apps
make build

# Run API in dev mode
make api-dev

# Run worker
make worker-run

# Build CLI
make cli-build

# Clean all artifacts
make clean

# See all available commands
make help
```

## Next Steps

1. Explore the individual READMEs in each app directory
2. Check out `shared/ARCHITECTURE.md` for system design
3. Extend the APIs with new endpoints
4. Add database integration
5. Deploy to your cloud provider

## Troubleshooting

### Node.js API won't start
```bash
cd api
npm install
npm run build
```

### Python Worker issues
```bash
cd worker
python3 -m venv venv
source venv/bin/activate
pip install --upgrade -r requirements.txt
```

### Go CLI build fails
```bash
cd cli
go mod tidy
go build main.go
```

## Need Help?

See individual README files:
- `api/README.md`
- `worker/README.md`
- `cli/README.md`
