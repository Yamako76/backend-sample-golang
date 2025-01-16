SHELL := /bin/sh

.PHONY: setup
setup:
	docker compose up -d --build

.PHONY: build
build:
	docker compose build

.PHONY: up
up:
	docker compose up -d

.PHONY: stop
stop:
	docker compose stop

.PHONY: down
down:
	docker compose down

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	go vet ./...
	staticcheck ./...

.PHONY: test
test: run
	go test ./...
