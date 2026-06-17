# API Service

TypeScript/Node.js REST API for task management.

## Setup

```bash
cd api
npm install
```

## Development

```bash
npm run dev
```

Server runs on `http://localhost:3000`

## Build

```bash
npm run build
```

## Endpoints

- `GET /health` - Health check
- `GET /tasks` - List all tasks
- `POST /tasks` - Create new task
- `GET /tasks/:id` - Get task by ID
- `PATCH /tasks/:id` - Update task
- `DELETE /tasks/:id` - Delete task

## Example Requests

```bash
# Create a task
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn TypeScript"}'

# List tasks
curl http://localhost:3000/tasks

# Update task
curl -X PATCH http://localhost:3000/tasks/1234 \
  -H "Content-Type: application/json" \
  -d '{"completed": true}'
```
