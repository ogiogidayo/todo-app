.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help
SHELL=/bin/bash
BIN_DIR := $(shell pwd)/bin
GOIMPORTS := $(abspath $(BIN_DIR)/goimports)

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t ogiogidayo/todo-app:${DOCKER_TAG} \
		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	go test -race -shuffle=on ./...

dry-migrate: ## Try migration
	mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo --dry-run < ./_tools/mysql/schema.sql

migrate:  ## Execute migration
	mysqldef -u todo -p todo -h 127.0.0.1 -P 33306 todo < ./_tools/mysql/schema.sql

generate: ## Generate codes
	go generate ./...
.PHONY: fmt

fmt: $(GOIMPORTS) ## format code
	find . -path ./proto -prune -o -name '*.go' -print | xargs $(GOIMPORTS) -w -local "github.com/ogiogidayo/todo-app"

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'