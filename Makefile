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

.PHONY: start-docker
start-docker:
ifeq ($(DETACH),true)
	@docker-compose up -d
else
	@docker-compose up
endif

.PHONY: stop-docker
stop-docker:
	@docker-compose down

.PHONY: logs
make logs:
	@docker-compose logs

.PHONY: start-app
start-app:
ifeq ($(DETACH),true)
	@. ./test.env && nohup go run main.go & bash ./script/wait-for-healthy-app.sh
else
	@. ./test.env && go run main.go
endif

.PHONY: test
test:
	@make start-docker DETACH=true
	@CONTAINER=mongo bash ./script/wait-for-healthy-container.sh
	@make start-app DETACH=true
	@go test ./... -count=1
	@make stop-docker
