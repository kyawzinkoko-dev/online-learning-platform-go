DB_USER=admin
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=learning_db
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: build run db-up db-down migrate-create migrate-up migrate-down help

## Help:
help: ## Show this help message with list of available commands
	@echo ""
	@echo "Available commands:"
	@echo "-------------------"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
	@echo ""

## App Management:
build: ## Compile and build the application binary
	@echo "Building the application..."
	@go build -o bin/api cmd/api/main.go

run: build ## Build and execute the application server
	@echo "Running the application..."
	@./bin/api

## Infrastructure Management:
db-up: ## Spin up PostgreSQL and Redis Docker containers
	@echo "Starting the database..."
	@docker compose up -d

db-down: ## Stop and remove all active project Docker containers
	@echo "Stopping the database..."
	@docker compose down

## Migration Management:
migrate-create: ## Create a new pair of migration files (Usage: make migrate-create name=filename)
	@echo "Creating a new migration..."
	@migrate create -ext sql -dir internal/database/migrations -seq $(name)

migrate-up: ## Run all pending database migrations (.up.sql files)
	@echo "Applying migrations..."
	@migrate -path internal/database/migrations -database "$(DB_URL)" up

migrate-down: ## Rollback the absolute last migration step (.down.sql file)
	@echo "Rolling back the last migration..."
	@migrate -path internal/database/migrations -database "$(DB_URL)" down 1

dev: ## Start active hot-reloading code server via Air
	@echo "Starting application with live-reload..."
	@air