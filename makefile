#* GET FILE ENV
include .env
export $(shell sed 's/=.*//' .env)

# * CONSTANTS
DOCKER_COMPOSE_DEV := docker-compose.dev.yml
DOCKER_COMPOSE_PRO := docker-compose.pro.yml

DOCKER_IMAGE_NAME :=nguyentientai/go-secure-auth-pro:lastest
DOCKERFILE_PATH := ./third_party/docker/go/Dockerfile

################# TODO: GOLANG #################
start:
	go run ./cmd/server/main.go

dev:
	go run fsnotify.go

################# TODO: DOCKER #################
build-pro:
	docker-compose -f $(DOCKER_COMPOSE_PRO) up -d --build

down-pro:
	docker-compose -f $(DOCKER_COMPOSE_PRO) down

build-dev:
	docker-compose -f $(DOCKER_COMPOSE_DEV) up -d --build

down-pro:
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


