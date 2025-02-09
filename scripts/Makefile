# Makefile

# Variables
APP_NAME=$(shell basename $(PWD))
SOURCE_FILE=main.go
GO=go

PROJECT_VERSION := $(shell git describe --abbrev=8 --always --dirty)
ifneq ($(CI_COMMIT_TAG),)
PROJECT_VERSION := $(CI_COMMIT_TAG)
endif

IMAGE_TAG=$(PROJECT_VERSION)
IMAGE_REPO=registry.docker.com/$(APP_NAME)
DOCKER_IMAGE=$(IMAGE_REPO):$(IMAGE_TAG)
DOCKER_IMAGE_LATEST=$(IMAGE_REPO):latest


all: build

.PHONY: build
build:
	$(GO) build -o ./bin/$(APP_NAME) $(SOURCE_FILE)

.PHONY: clean
clean:
	@rm -f ./bin/$(APP_NAME)

.PHONY: run
run: build
	@./bin/$(APP_NAME)

.PHONY: docker-build
docker-build:
	@docker build --platform linux/amd64 -t $(DOCKER_IMAGE) .

.PHONY: docker-push
docker-push:
	@docker push $(DOCKER_IMAGE)

.PHONY: docker-push-latest
docker-push-latest:
	@docker tag $(DOCKER_IMAGE) $(DOCKER_IMAGE_LATEST)
	@docker push $(DOCKER_IMAGE_LATEST)

.PHONY: all

.PHONY: version
version: ## Display version.
	@echo $(PROJECT_VERSION)