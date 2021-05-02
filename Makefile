.PHONY: run-lint

default: build

lint:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin latest
	golangci-lint run --fix

run:
	go run ./cmd/console/main.go

run-lint: lint
	go run ./cmd/console/main.go

vendor:
	go mod vendor

build:
	go build ./cmd/console/main.go

tidy:
	go mod tidy

tcp-server:
	go run ./cmd/tcp/server.go

tcp-client-console:
	go run ./cmd/tcp/client_console.go

tcp-client-csv:
	go run ./cmd/tcp/client_csv.go

http-server:
	go run ./cmd/server/server_http.go

run-mysql:
	go run ./cmd/database/mysql.go

run-postgres:
	go run ./cmd/database/postgres.go
