#* GET FILE ENV
include .env
export $(shell sed 's/=.*//' .env)

# * FILE RUN GO
GO_SERVER_PRO := ./cmd/server/main.go
GO_SERVER_DEV:= ./fsnotify.go
GO_SERVER_CRON := ./cmd/cronjob/main.go

# * DOCKER COMPOSE
DOCKER_COMPOSE_DEV := docker-compose.dev.yml
DOCKER_COMPOSE_PRO := docker-compose.pro.yml

# * DOCKER HUB
SERVER_IMAGE_NAME :=nguyentientai/go-secure-auth-pro:lastest
CRON_IMAGE_NAME :=nguyentientai/go_cronjob_auth:lastest

# * DOCKER FILE
DOCKER_FILE_PATH := ./third_party/docker/go/Dockerfile

#* DOCKER CONTAINER
CONTAINER_SERVICE_AUTH := service_auth

# * DOCKER IMAGE
TARGET_SERVER := server
TARGET_CRON := cron

# * FOLDER
SWAGGER_DIR=./docs/swagger

################# TODO: GOLANG #################
start:
	go run $(GO_SERVER_PRO)

dev:
	go run $(GO_SERVER_DEV)

cron:
	go run $(GO_SERVER_CRON)

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
# Build and tag the server image
server-image-tag:
	docker build --target $(TARGET_SERVER) -t $(SERVER_IMAGE_NAME) -f $(DOCKER_FILE_PATH) .

# Build and tag the cron image
cron-image-tag:
	docker build --target $(TARGET_CRON) -t $(CRON_IMAGE_NAME) -f $(DOCKER_FILE_PATH) .


# Push the server image to the registry
push-server: server-image-tag
	docker push $(SERVER_IMAGE_NAME)

# Push the cron image to the registry
push-cron: cron-image-tag
	docker push $(CRON_IMAGE_NAME)

# Combined target to build and push both images
build-and-push-all: push-server push-cron
	@echo "Both server and cron images have been built and pushed successfully."

################# TODO: SQLC #################
sqlc:
	sqlc generate

################# TODO: SWAGGER #################
swagger:
	@echo "Generating Swagger documentation..."
	swag init --parseDependency -g  $(GO_SERVER_PRO) -o $(SWAGGER_DIR)




