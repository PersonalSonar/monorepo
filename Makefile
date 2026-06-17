.PHONY: help install build run test clean setup

help:
	@echo "Monorepo Commands:"
	@echo "=================="
	@echo "make setup       - Install all dependencies"
	@echo "make install     - Install dependencies for all apps"
	@echo "make build       - Build all apps"
	@echo "make run         - Run all apps (requires tmux or manual terminals)"
	@echo "make test        - Test all apps"
	@echo "make clean       - Clean all build artifacts"
	@echo ""
	@echo "Individual app commands:"
	@echo "make api-dev     - Run API in development mode"
	@echo "make api-build   - Build API"
	@echo "make worker-run  - Run worker service"
	@echo "make cli-build   - Build CLI tool"

setup: install

install:
	cd api && npm install
	cd worker && pip install -r requirements.txt
	cd cli && go mod download

build: build-api build-worker build-cli

build-api:
	cd api && npm run build

build-worker:
	@echo "Worker is a Python script, no build needed"

build-cli:
	cd cli && go build -o monorepo main.go

run:
	@echo "Starting all services..."
	@echo "Run these in separate terminals:"
	@echo ""
	@echo "Terminal 1 - API:"
	@echo "  cd api && npm run dev"
	@echo ""
	@echo "Terminal 2 - Worker:"
	@echo "  cd worker && python worker.py"
	@echo ""
	@echo "Terminal 3 - CLI:"
	@echo "  cd cli && go run main.go status"

test:
	cd api && npm test || true
	cd worker && python -m pytest || true

clean:
	rm -rf api/dist api/node_modules
	rm -rf worker/__pycache__ worker/jobs.json
	rm -f cli/monorepo

# Individual app targets
api-dev:
	cd api && npm run dev

api-build:
	cd api && npm run build

worker-run:
	cd worker && python worker.py

cli-build:
	cd cli && go build -o monorepo main.go

cli-run:
	cd cli && ./monorepo
