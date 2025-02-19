include .env

build: 
	@go build -o ./bin/auth ./cmd/auth/main.go

run: build
	@./bin/auth

migrate_up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${PG_DSN} goose -dir=./migrations up

migrate_down:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${PG_DSN} goose -dir=./migrations down
