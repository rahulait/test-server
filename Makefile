# Configurable variables
APP_NAME ?= test-server
VERSION  ?= latest
DOCKER_USER ?= rahulait
IMAGE_NAME ?= $(DOCKER_USER)/$(APP_NAME):$(VERSION)

# Build the Go binary
build:
	go build -o $(APP_NAME)

# Build the Docker image
docker-build: build
	docker build -t $(IMAGE_NAME) .

# Push the Docker image to Docker Hub
docker-push:
	docker push $(IMAGE_NAME)

# Clean up binary
clean:
	rm -f $(APP_NAME)

# Full pipeline: build binary, docker image, and push
release: docker-build docker-push

.PHONY: build docker-build docker-push clean release
