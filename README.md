# Multi-Language Monorepo

A monorepo containing three distinct applications in different languages.

## Structure

```
.
├── api/           - TypeScript/Node.js REST API
├── worker/        - Python data processing service
├── cli/           - Go command-line tool
└── shared/        - Shared resources and documentation
```

## Apps

### API (TypeScript/Node.js)
Express.js REST API with database models. Port: 3000

### Worker (Python)
Background job processor using asyncio. Processes tasks from queue.

### CLI (Go)
Command-line tool for managing and monitoring the monorepo services.

## Development

Each app is self-contained with its own dependencies and build process. See individual README files in each app directory.
