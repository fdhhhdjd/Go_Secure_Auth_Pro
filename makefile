#* GET FILE ENV
include .env
export $(shell sed 's/=.*//' .env)

# * CONSTANTS
DOCKER_COMPOSE := docker-compose.yml

################# TODO: GOLANG #################
start:
	go run ./cmd/server/main.go

dev:
	go run fsnotify.go

################# TODO: DOCKER #################
build:
	docker-compose -f $(DOCKER_COMPOSE) up -d --build

################# TODO: SQLC #################
sqlc:
	sqlc generate


