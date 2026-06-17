# CLI Tool

Go command-line utility for managing monorepo services.

## Setup

```bash
cd cli
go mod download
```

## Build

```bash
go build -o monorepo main.go
```

## Usage

```bash
./monorepo [command]
```

## Available Commands

### status
Check status of all services:
```bash
./monorepo status
```

### list
List all apps in the monorepo:
```bash
./monorepo list
```

### logs
View logs from a specific service:
```bash
./monorepo logs api
./monorepo logs worker
```

### info
Show detailed monorepo information:
```bash
./monorepo info
```

## Installation

To install globally:
```bash
go build -o /usr/local/bin/monorepo main.go
```

Then use anywhere:
```bash
monorepo status
```

## Development

```bash
go run main.go [command]
```

## Dependencies

- `cobra` - CLI framework
- `viper` - Configuration management
