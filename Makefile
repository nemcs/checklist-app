.PHONY: up down build clean test lint

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build

clean:
	docker-compose prune -f

test:
	go test ./...

lint:
	golangci-lint run
