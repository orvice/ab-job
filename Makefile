.DEFAULT_GOAL := build

APP_NAME=ab-job
APP_CMD_DIR=cmd/$(APP_NAME)
APP_BINARY=bin/$(APP_NAME)
APP_BINARY_UNIX=bin/$(APP_NAME)_unix_amd64
DOCKER_IMAGE=orvice/ab-job

all: build

.PHONY: test
test: ## test
	go test -v ./...

.PHONY: build
build: ## build
	CGO_ENABLED=0 go build -o $(APP_BINARY) -v $(APP_CMD_DIR)/main.go

docker:
	docker build -t $DOCKER_IMAGE .