.PHONY: help dev dev-build start stop logs clean backend-dev frontend-dev

help: ## Show this help message
	@echo "Sanctor Monorepo - Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

dev: ## Start development environment with hot reload
	docker-compose -f docker-compose.dev.yml up

dev-build: ## Build and start development environment
	docker-compose -f docker-compose.dev.yml up --build

start: ## Start production environment
	docker-compose up

build: ## Build production images
	docker-compose build

stop: ## Stop all containers
	docker-compose down

logs: ## Show container logs
	docker-compose logs -f

clean: ## Remove all containers, images, and volumes
	docker-compose down -v --rmi all

backend-dev: ## Run backend locally (requires Go)
	cd apps/backend && go run main.go

frontend-dev: ## Run frontend locally (requires Node.js)
	cd apps/frontend && npm start

frontend-install: ## Install frontend dependencies
	cd apps/frontend && npm install

test-backend: ## Run backend tests
	cd apps/backend && go test ./...

test-frontend: ## Run frontend tests
	cd apps/frontend && npm test
