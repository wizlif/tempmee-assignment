# Go Take-Home Assignment

This is a Go project that uses gRPC-gateway and SQLite. It uses Viper for managing application configurations which are stored in an `app.env` file. The development database is `dev.db`.

## Prerequisites

- Go
- SQLite
- gRPC
- [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/)
- [Viper](https://github.com/spf13/viper)
- [Buf](https://buf.build/docs/introduction)
- [sqlc](https://sqlc.dev/docs/introduction/)
- [VSCode REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

## Database Documents

The database documents are available at [dbdocs.io](https://dbdocs.io/wizlif.144/tempmee).

Password

```
'zs&U+sI"qtC>R]&y.4B:Y&QhP3]fx08i{5])z$k7so,(sy(4$/L8_Dg[37Z>tP
```

## Makefile

A Makefile is provided to help with common tasks. Here are the commands you can run:

- `make create_migration` :              Create database migration e.g make create_migration name=initial_schema
- `make db_schema`:                      Build SQL from dbml
- `make migratedown`:                    Run migrations down
- `make migrateup`:                      Run migrations up
- `make mock`:                           Generate mocks
- `make prod`:                           Run project with docker compose
- `make proto`:                          Build proto files
- `make server`:                         Run dev server
- `make sqlc`:                           Build sqlc database files
- `make test`:                           Run tests
- `make test_coverage`:                  Run test with gocov coverage

## sqlc for Database Migrations

sqlc is used to create database migrations. To create a new migration, use the command `make create_migration name=<MIGRATION NAME>`. Replace `<MIGRATION NAME>` with the name of your migration.

## Buf for Protocol Buffers

We use Buf for managing our Protocol Buffers. You can use the command `make proto` to lint, format, and generate protobuf stubs and Swagger documentation via `protoc-gen-openapiv2`.

We also use `protovalidate` for generating validations for our protobuf models.

## REST Client Extension for VS Code

We use the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension for VS Code for executing HTTP requests. You can find our test HTTP requests in the [requests.http](./requests.http) file.

# Getting Started

1. Clone the repository
2. Run `make server` to start the server or with docker use `make prod`

For more information, refer to the Makefile.

# Further Reading

- [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/)
- [SQLite](https://sqlite.org/index.html)
- [Viper](https://github.com/spf13/viper)
- [Buf](https://buf.build/docs/introduction)
- [sqlc](https://sqlc.dev/docs/introduction/)
- [VSCode REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

