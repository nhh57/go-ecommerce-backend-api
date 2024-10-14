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

.PHONY: run

.PHONE:air

