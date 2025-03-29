.PHONY: setup run build clean test docker-up docker-down generate

# Default target
all: setup generate build

# Setup the project
setup:
	go mod download

# Generate GraphQL code
generate:
	go run github.com/99designs/gqlgen generate

# Build the application
build:
	go build -o bin/mytolet-api ./cmd/server

# Run the application
run:
	go run ./cmd/server

# Run tests
test:
	go test ./...

# Clean binaries
clean:
	rm -rf bin/

# Start Docker containers
docker-up:
	docker-compose up -d

# Stop Docker containers
docker-down:
	docker-compose down

# Start Docker containers with build
docker-build:
	docker-compose up -d --build

# View logs
docker-logs:
	docker-compose logs -f

# Database migrations
db-migrate:
	go run ./cmd/migrate

# Seed database with test data
db-seed:
	go run ./cmd/seed