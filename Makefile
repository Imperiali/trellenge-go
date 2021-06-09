CONTAINER_NAME=trellenge-go
DOCKER_TARGET?=development
PWD=$(shell pwd)
PORT=8080

build-development-image:
	@docker build --target development -t $(CONTAINER_NAME):development .

build-production-image:
	@docker build --target production -t $(CONTAINER_NAME):production .

check-development-image:
ifeq ($(shell docker images -q $(CONTAINER_NAME) 2> /dev/null | wc -l), 0)
	@echo "Image not found, creating a new one"
	@make build-development-image
endif

start-dev: check-development-image
	@echo "Starting challenge as development mode"
	@docker run --name $(CONTAINER_NAME) --env-file .env -v $(PWD):/app -p $(PORT):80 --rm $(CONTAINER_NAME):development

start-prod: build-production-image
	@echo "Starting challenge as production mode"
	@docker run --name $(CONTAINER_NAME) --env-file .env -p $(PORT):8080 --rm $(CONTAINER_NAME):production

test:
	@echo "Running tests..."
	@go test ./...
