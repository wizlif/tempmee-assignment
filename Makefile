SHELL := /usr/bin/env bash

PROTO_DIR    := proto

PROTO_FILES  = $(shell find $(PROTO_DIR)/ -type f -name '*.proto')
OPENAPI_PROTO_FILES  = $(shell find $(PROTO_DIR)/protoc-gen-openapiv2 -type f -name '*.proto')
PROTO_OUT_DIR   := pb

DOC_DIR   := doc

DB_URL=sqlite://dev.db

proto: ## Build proto files
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	buf lint
	buf generate
	statik -src=$(DOC_DIR)/swagger -dest=./doc

db_schema: ## Build SQL from dbml
	dbml2sql --sqlite -o doc/schema.sql doc/db.dbml

sqlc: ## Build sqlc database files
	sqlc generate

mock: ## Generate mocks
	mockgen -package mockdb -destination db/mock/store.go github.com/wizlif/tempmee_assignment/db/sqlc Store

server: ## Run dev server
	go run main.go

migrateup: ## Run migrations up
	migrate --path db/migration --database "$(DB_URL)" --verbose up

migrateup1: ## Run migrations single step up
	migrate --path db/migration --database "$(DB_URL)" --verbose up 1

migratedown: ## Run migrations down
	migrate --path db/migration --database "$(DB_URL)" --verbose down

migratedown1: ## Run migrations single step down
	migrate --path db/migration --database "$(DB_URL)" --verbose down 1

test: ## Run tests
	go test -v -cover -short ./...

test_coverage: ## Run test with gocov coverage
	go test -coverprofile=coverage.out ./...
	gocov convert coverage.out | gocov report

create_migration: ## Create database migration e.g make create_migration name=initial_schema
	migrate create -ext sql -dir db/migration -seq $(name)


help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help proto create_migration test migratedown migratedown1 migrateup migrateup1 mock sqlc db_schema proto test_coverage
.DEFAULT_GOAL := help
