GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= "root:Admin123@tcp(localhost:3306)/shopDev"
GOOSE_MIGRATION_DIR ?= sql/schema
# name app
APP_NAME := server

docker_build:
	docker-compose up -d --build
	docker-compose ps

docker_stop:
	docker-compose down

dev:
	go run cmd/$(APP_NAME)/main.go

docker_up:
	docker-compose up -d

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen:
	sqlc generate

.PHONY: dev downse upse resetse docker_build docker_stop docker_up

.PHONE: air

