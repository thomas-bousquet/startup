.PHONY: docker-login
docker-login:
	@docker login --username=$(DOCKER_USERNAME) --password=$(DOCKER_ACCESS_TOKEN)

.PHONY: docker-build
docker-build:
	@docker build -t $(IMAGE_NAME) --build-arg APP_VERSION=$(APP_VERSION) .

.PHONY: docker-tag
docker-tag:
	@docker tag $(IMAGE_NAME) $(DOCKER_USERNAME)/$(DOCKER_REPOSITORY):$(TAG)
	@docker tag $(IMAGE_NAME) $(DOCKER_USERNAME)/$(DOCKER_REPOSITORY):latest

.PHONY: docker-push
docker-push:
	@docker push $(DOCKER_USERNAME)/$(DOCKER_REPOSITORY):$(TAG)
	@docker push $(DOCKER_USERNAME)/$(DOCKER_REPOSITORY):latest

.PHONY: start
start:
ifeq ($(DETACH),true)
	@docker-compose up -d
else
	@docker-compose up
endif

.PHONY: stop
stop:
	@docker-compose down

.PHONY: test
test:
	@make start DETACH=true
	@CONTAINER=mongo bash ./script/wait-for-healthy-container.sh
	@CONTAINER=user-service bash ./script/wait-for-healthy-container.sh
	@go test ./... -count=1
	@make stop
