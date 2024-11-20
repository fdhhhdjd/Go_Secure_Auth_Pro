################# TODO: CONSTANTS #################

#* GET FILE ENV
include .env
export $(shell sed 's/=.*//' .env)

# * FILE RUN GO
GO_SERVER_PRO := ./cmd/server/main.go
GO_SERVER_DEV:= ./fsnotify.go
GO_SERVER_CRON := ./cmd/cronjob/main.go
GO_CONSUMER := ./cmd/queue/main.go

# * DOCKER COMPOSE
DOCKER_COMPOSE_DEV := docker-compose.dev.yml
DOCKER_COMPOSE_PRO := docker-compose.pro.yml

# * DOCKER HUB
SERVER_IMAGE_NAME := nguyentientai/go-secure-auth-pro:lastest
CRON_IMAGE_NAME := nguyentientai/go_cronjob_auth:lastest
QUEUE_IMAGE_NAME := nguyentientai/go_message_queue_auth:lastest


# * DOCKER FILE
DOCKER_FILE_PATH := ./third_party/docker/go/Dockerfile
DOCKER_FILE_CRON_PATH := ./third_party/docker/go/Dockerfile-cron
DOCKER_FILE_QUEUE_PATH := ./third_party/docker/go/Dockerfile-queue

#* DOCKER CONTAINER
CONTAINER_SERVICE_AUTH := service_auth
CONTAINER_SERVICE_CRON := service_cron
CONTAINER_SERVICE_QUEUE := service_queue


# * FOLDER
SWAGGER_DIR=./docs/swagger

################# TODO: GO #################
start:
	go run $(GO_SERVER_PRO)

dev:
	go run $(GO_SERVER_DEV)

air:
	air

cron:
	go run $(GO_SERVER_CRON)

consumer:
	go run $(GO_CONSUMER)
    
################# TODO: DOCKER #################
build-pro:
	docker-compose -f $(DOCKER_COMPOSE_PRO) up -d --build

down-pro:
	docker-compose -f $(DOCKER_COMPOSE_PRO) down

build-dev:
	docker-compose -f $(DOCKER_COMPOSE_DEV) up -d --build

down-dev:
	docker-compose -f $(DOCKER_COMPOSE_DEV) down

update-server:
	docker-compose -f $(DOCKER_COMPOSE_PRO) pull $(CONTAINER_SERVICE_AUTH)
	docker-compose -f $(DOCKER_COMPOSE_PRO) up -d --no-deps $(CONTAINER_SERVICE_AUTH)

update-cron:
	docker-compose -f $(DOCKER_COMPOSE_PRO) pull $(CONTAINER_SERVICE_QUEUE)
	docker-compose -f $(DOCKER_COMPOSE_PRO) up -d --no-deps $(CONTAINER_SERVICE_QUEUE)

update-image: update-server update-cron
	@echo "Both server and cron images updated successfully."

################# TODO: DOCKER HUB #################
# Build and tag the server image
server-image-tag:
	docker build -t $(SERVER_IMAGE_NAME) -f $(DOCKER_FILE_PATH) .

# Build and tag the cron image
cron-image-tag:
	docker build -t $(CRON_IMAGE_NAME) -f $(DOCKER_FILE_CRON_PATH) .

# Build and tag the cron image
queue-image-tag:
	docker build -t $(QUEUE_IMAGE_NAME) -f $(DOCKER_FILE_QUEUE_PATH) .

# Push the server image to the registry
push-server: server-image-tag
	docker push $(SERVER_IMAGE_NAME)

# Push the cron image to the registry
push-cron: cron-image-tag
	docker push $(CRON_IMAGE_NAME)

# Push the cron image to the registry
push-queue: queue-image-tag
	docker push $(QUEUE_IMAGE_NAME)

# Combined target to build and push both images
build-and-push-all: push-server push-cron
	@echo "Both server and cron images have been built and pushed successfully."

################# TODO: SQLC #################
# Generate SQLC code
sqlc:
	sqlc generate

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	swag init --parseDependency -g $(GO_SERVER_PRO) -o $(SWAGGER_DIR)




