.PHONY: docker-up docker-down build integration-test

docker-up:
	@go mod vendor
	@docker-compose up

docker-down:
	@docker-compose down

build:
	docker build -t startup .

integration-test:
	@go test -v ./it-test -count=1

docker-logs:
	docker-compose logs