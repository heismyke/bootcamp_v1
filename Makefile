# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go
# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi
# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

createdb: 
	docker exec -it postgres17 createdb --username=root --owner=root devcamper

dropdb: 
	docker exec -it postgres17 dropdb devcamper

migrateup:
	migrate -path ./internal/database/migration -database "postgres://root:MikeSera2022@localhost:5432/devcamper?sslmode=disable" -verbose up

migratedown:
	migrate -path ./internal/database/migration -database "postgres://root:MikeSera2022@localhost:5432/devcamper?sslmode=disable" -verbose down

migrateforce:
	@read -p "Enter version to force (0 to reset): " version; \
	migrate -path ./internal/database/migration -database "postgres://root:MikeSera2022@localhost:5432/devcamper?sslmode=disable" force $$version

generate:
	@read -p "Are you sure you want to generate queries? (yes/no): " response; \
	if [ $$response = "yes" ]; then \
		echo "Generating queries..."; \
		sqlc generate; \
	elif [ $$response = "no" ]; then \
		echo "Exiting..."; \
		exit 0; \
	else \
		echo "Invalid input, exiting..."; \
		exit 1; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@@go test ./internal/database/sqlc -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch docker-run docker-down itest createdb dropdb migrateup migratedown migrateforce generate
