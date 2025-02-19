include .env

build:
	@go build -o ./bin/favorite cmd/favorite/main.go

run: build
	@./bin/favorite

migrate_up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${PG_DSN} goose -dir=./migrations up

migrate_down:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=${PG_DSN} goose -dir=./migrations down
