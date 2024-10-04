-include .env
DEFAULT_GOAL := local

.PHONY: run
run:
	go run cmd/test/*

.PHONY: build
build:
	go build -o srv cmd/test/*

.PHONY: migrate
migrate:
	goose -dir="./db/migrations" postgres "host=localhost port=$(PG_PORT) user=$(PG_USER) password=$(PG_PWD) database=$(PG_DB_NAME) sslmode=disable" status -v
	goose -dir="./db/migrations" postgres "host=localhost port=$(PG_PORT) user=$(PG_USER) password=$(PG_PWD) database=$(PG_DB_NAME) sslmode=disable" up -v

.PHONY: new-migration
new-migration:
	goose -dir="./db/migrations" create $(name) sql

.PHONY: test
test:
	go test -v ./...

.PHONY: test-race
test-race:
	go test -v -short -race ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml ./...

.PHONY: docker-build
docker-build:
	docker compose --file docker/docker-compose.local.yml --env-file docker/local.env down
	docker compose --file docker/docker-compose.local.yml -p "test"  --env-file docker/local.env up --detach

.PHONY: docker-start
docker-start:
	docker compose --file docker/docker-compose.local.yml --env-file docker/local.env build --pull
	docker compose --file docker/docker-compose.local.yml -p "test"  --env-file docker/local.env up --detach

.PHONY: docker-stop
docker-stop:
	docker compose --file docker/docker-compose.local.yml --env-file docker/local.env down --volumes
	docker compose --file docker/docker-compose.local.yml --env-file docker/local.env rm --force

.PHONY: docker-logs
docker-logs:
	docker compose --file docker/docker-compose.local.yml --env-file docker/local.env logs

.PHONY: rebuild
rebuild:
	make docker-stop || true
	make docker-start

.PHONY: generate
generate:
	cd api/test && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto