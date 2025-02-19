include .env

build:
	@go build -o bin/pet cmd/pet/main.go

run: build
	@./bin/pet

test:
	go test ./...

test_100:
	go test ./... -count 100

migrate_up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${PG_DSN} goose -dir=./migrations up

migrate_down:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${PG_DSN} goose -dir=./migrations down
