SHELL := /usr/bin/env bash

PROTO_DIR    := proto

PROTO_FILES  = $(shell find $(PROTO_DIR)/ -type f -name '*.proto')
OPENAPI_PROTO_FILES  = $(shell find $(PROTO_DIR)/protoc-gen-openapiv2 -type f -name '*.proto')
PROTO_OUT_DIR   := pb

DOC_DIR   := doc

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

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help proto
.DEFAULT_GOAL := help
