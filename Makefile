.PHONY: docker-up docker-down build integration-test docker-logs

docker-up:
	@go mod vendor
	@docker-compose up

docker-down:
	@docker-compose down

build:
	@go build

integration-test:
	@go test -v ./it-test -count=1

docker-logs:
	docker-compose logs