DB_DRIVER=postgres
MIGRATIONS_DIR=./migrations

include .env
export

GOOSE=go run github.com/pressly/goose/v3/cmd/goose

migrate-up:
	@echo "Applying all up migrations"
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DATABASE_URL)" up

migrate-down:
	@echo "Reverting last migration"
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DATABASE_URL)" down

migrate-status:
	@echo "Checking migration status"
	@$(GOOSE) -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DATABASE_URL)" status

migrate-create:
	@echo "Creating new migration file"
	@$(GOOSE) -dir $(MIGRATIONS_DIR) create $(name) sql

.PHONY: migrate-up migrate-down migrate-status migrate-create

