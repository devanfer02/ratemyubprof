# variables
include .db.env
DB_URL = "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

.PHONY: help
help:
	@echo "Choose a command:"
	@echo "  make run            				- Run the Go application"
	@echo "  make air            				- Run the Go application with live reloading using air"
	@echo "  make migrate           			- Run migration"
	@echo "  make seed              			- Run seeder to fill data in database"
	@echo "  make test           				- Run Go unit tests defined in ./tests"
	@echo "  make migrate-up     				- Apply all migrations with migrate-cli"
	@echo "  make migrate-down   				- Rollback all migrations"
	@echo "  make create-migration name=<name> 		- Create a new migration file"

.PHONY: run
run:
	go run ./cmd/main/main.go $(args)

.PHONY: seed 
seed:
	go run ./cmd/main/main.go seed

.PHONY: migrate 
migrate:
	go run ./cmd/main/main.go migrate

.PHONY: test
test:
	go test ./tests/...

.PHONY: air
air:
	air -c .air.toml

.PHONY: compose-up
compose-up:
ifndef file
	docker compose up -d
else
	docker compose -f $(file) up -d
endif 

.PHONY: compose-down
compose-down:
ifndef file
	docker compose down
else
	docker compose -f $(file) down 
endif 

.PHONY: test
test:
	go test ./tests/...

.PHONY: migrate-up
migrate-up:
	migrate -path ./internal/infra/database/postgres/migrations -database ${DB_URL} up

.PHONY: migrate-down
migrate-down:
ifndef version
	migrate -path ./internal/infra/database/postgres/migrations -database ${DB_URL} down 
else 
	migrate -path ./internal/infra/database/postgres/migrations -database ${DB_URL} down $(version)
endif 

.PHONY: migrate-force
migrate-force:
ifndef version 
	$(error "Migration version not specified. Use 'make migrate-force version=<version>'")
endif 
	migrate -path ./internal/infra/database/postgres/migrations -database ${DB_URL} force $(version)

.PHONY: migrate-create
migrate-create:
ifndef name 
	$(error "Migration name not specified. Use 'make create-migration name=<name>'")
endif 
	migrate create -ext sql -dir ./internal/infra/database/postgres/migrations $(name)