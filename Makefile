.PHONY: start-local
start-local:
	@docker-compose -f docker-compose.local.yml up --remove-orphans --force-recreate --build

.PHONY: stop-local
stop-local:
	@docker-compose -f docker-compose.local.yml down

.PHONY: docker-login
docker-login:
	@docker login --username=$(DOCKER_USERNAME) --password=$(DOCKER_ACCESS_TOKEN)

.PHONY: docker-build
docker-build:
	@docker build -t $(IMAGE_NAME) .

.PHONY: docker-tag
docker-tag:
	@docker tag $(IMAGE_NAME) $(DOCKER_USERNAME)/$(DOCKER_REPOSITORY):latest

.PHONY: docker-push
docker-push:
	@docker push $(DOCKER_USERNAME)/$(DOCKER_REPOSITORY):latest

.PHONY: start
start:
ifeq ($(DETACH),true)
	@docker-compose up --remove-orphans --force-recreate --build -d
else
	@docker-compose up --remove-orphans --force-recreate --build
endif

.PHONY: stop
stop:
	@docker-compose down

.PHONY: logs
make logs:
	@docker-compose logs

.PHONY: test
test:
	@go test ./... -count=1
