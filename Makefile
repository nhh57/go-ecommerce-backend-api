GOOSE_DRIVER ?= "mysql"
GOOSE_DBSTRING= "root:Admin123@tcp(localhost:3306)/shopDEV"
GOOSE_MIGRATION_DIR ?= sql/schema

# name app
APP_NAME := server


dev:
	go run cmd/$(APP_NAME)/main.go
run:
	docker compose up -d && go run cmd/$(APP_NAME)

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: run downse upse resetse

.PHONE: air

