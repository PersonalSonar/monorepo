# Monorepo Architecture

## Overview

This is a multi-language monorepo containing three independent services that work together:

```
┌─────────────────────────────────────┐
│         CLI Tool (Go)               │
│     Management & Monitoring         │
└────────────────┬────────────────────┘
                 │
        ┌────────┴────────┐
        │                 │
┌───────▼────────┐  ┌─────▼──────────┐
│  API (Node.js) │  │ Worker (Python)│
│  REST Backend  │  │ Job Processing │
└────────────────┘  └────────────────┘
```

## Services

### API Service (TypeScript/Node.js)
- **Language**: TypeScript compiled to JavaScript
- **Framework**: Express.js
- **Port**: 3000
- **Purpose**: Provides REST endpoints for task management
- **Location**: `/api`

**Key Endpoints**:
- `GET /health` - Service health check
- `GET /tasks` - Retrieve all tasks
- `POST /tasks` - Create new task
- `PATCH /tasks/:id` - Update task
- `DELETE /tasks/:id` - Delete task

### Worker Service (Python)
- **Language**: Python 3.8+
- **Framework**: asyncio
- **Purpose**: Background job processing and data handling
- **Location**: `/worker`

**Features**:
- Async job queue
- Job persistence
- Status tracking (pending → processing → completed)
- Scalable architecture

### CLI Tool (Go)
- **Language**: Go 1.21+
- **Framework**: Cobra (CLI framework)
- **Purpose**: Command-line interface for service management
- **Location**: `/cli`

**Commands**:
- `status` - Check all services
- `list` - List all apps
- `logs` - View service logs
- `info` - Show monorepo information

## Communication

- **API to CLI**: HTTP requests to localhost:3000
- **Worker to API**: (Can be extended via HTTP calls)
- **CLI to Services**: Direct HTTP/Process calls

## Data Flow

1. User submits task via API
2. API creates task and returns ID
3. Worker picks up task from queue
4. Worker processes and updates status
5. CLI can query status of both services

## Development Workflow

```
1. Install dependencies:      make install
2. Build all services:        make build
3. Run individual services:   make api-dev, make worker-run, make cli-build
4. Test changes:              make test
5. Clean artifacts:           make clean
```

## Deployment Considerations

- Each service can be deployed independently
- Services are loosely coupled via HTTP/Queues
- Horizontal scaling possible for API and Worker
- CLI can run on any machine with access to services

## Future Enhancements

- Add database layer (PostgreSQL/MongoDB)
- Implement inter-service message queue (Redis/RabbitMQ)
- Add authentication/authorization
- Containerize with Docker
- Add monitoring/metrics (Prometheus)
- Add logging aggregation (ELK Stack)
