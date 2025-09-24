include .env
export $(shell sed 's/=.*//' .env)

migrate-up:
	@DATABASE_URL=$(DATABASE_URL) go run ./cmd/migrate up

migrate-reset:
	@DATABASE_URL=$(DATABASE_URL) go run ./cmd/migrate reset

migrate-status:
	@DATABASE_URL=$(DATABASE_URL) go run ./cmd/migrate status

db-up:
	docker compose up -d

db-down:
	docker compose down

db-restart: db-down db-up

psql:
	docker compose exec db psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)
