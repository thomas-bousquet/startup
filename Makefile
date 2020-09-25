.PHONY: start-dev stop-dev build-docker start stop-prod test build-test-app

start-dev:
	@docker-compose -f docker-compose.dev.yml up

stop-dev:
	@docker-compose -f docker-compose.dev.yml down

build-docker:
	@docker build -t startup .

start:
	@docker-compose -f docker-compose.test.yml up --remove-orphans --force-recreate --build

stop:
	@docker-compose -f docker-compose.test.yml down

test:
	@go test ./... -count=1
