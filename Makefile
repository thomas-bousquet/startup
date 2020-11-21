.PHONY: start-dev stop-dev build-docker start stop-prod test build-test-app

start-dev:
	@docker-compose -f docker-compose.dev.yml up

stop-dev:
	@docker-compose -f docker-compose.dev.yml down

docker-login:
	@docker login --username=$(DOCKER_USERNAME) --password=$(DOCKER_ACCESS_TOKEN)

docker-build:
	@docker build -t $(IMAGE_NAME) .

docker-tag:
	@docker tag $(IMAGE_NAME) $(DOCKER_USERNAME)/$(DOCKER_REPOSITORY):latest

docker-push:
	@docker push $(DOCKER_USERNAME)/$(DOCKER_REPOSITORY):latest

start:
	@docker-compose -f docker-compose.test.yml up --remove-orphans --force-recreate --build

stop:
	@docker-compose -f docker-compose.test.yml down

test:
	@go test ./... -count=1
