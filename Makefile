.PHONY: docker-up docker-down integration-test

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

build:
	go build

integration-test: #docker-up
	go test -v ./it-test
	#make docker-down