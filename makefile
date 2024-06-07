#* GET FILE ENV
include .env
export $(shell sed 's/=.*//' .env)

# * CONSTANTS
GO_SERVER_PRO := ./cmd/server/main.go
GO_SERVER_DEV:= ./fsnotify.go

DOCKER_COMPOSE_DEV := docker-compose.dev.yml
DOCKER_COMPOSE_PRO := docker-compose.pro.yml

DOCKER_IMAGE_NAME :=nguyentientai/go-secure-auth-pro:lastest
DOCKERFILE_PATH := ./third_party/docker/go/Dockerfile

SWAGGER_DIR=./docs/swagger

CONTAINER_SERVICE_AUTH := service_auth


################# TODO: GOLANG #################
start:
	go run $(GO_SERVER_PRO)

dev:
	go run $(GO_SERVER_DEV)

################# TODO: DOCKER #################
build-pro:
	docker-compose -f $(DOCKER_COMPOSE_PRO) up -d --build

down-pro:
	docker-compose -f $(DOCKER_COMPOSE_PRO) down

update-image:
	docker-compose -f $(DOCKER_COMPOSE_PRO) pull $(CONTAINER_SERVICE_AUTH)
	docker-compose -f $(DOCKER_COMPOSE_PRO) up -d --no-deps $(CONTAINER_SERVICE_AUTH)

build-dev:
	docker-compose -f $(DOCKER_COMPOSE_DEV) up -d --build

down-dev:
	docker-compose -f $(DOCKER_COMPOSE_DEV) down


################# TODO: DOCKER HUB #################
image-tag:
	docker build -t $(DOCKER_IMAGE_NAME) -f $(DOCKERFILE_PATH) .

push-registry:
	docker push $(DOCKER_IMAGE_NAME)

update-registry:
	make image-tag
	make push-registry

################# TODO: SQLC #################
sqlc:
	sqlc generate

################# TODO: SWAGGER #################
swagger:
	@echo "Generating Swagger documentation..."
	swag init --parseDependency -g  $(GO_SERVER_PRO) -o $(SWAGGER_DIR)




