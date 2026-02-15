# Database Config
DB_URL=postgres://admin:password@localhost:5432/wyd_db?sslmode=disable

# Migration path
MIGRATION_PATH=pkg/storage/postgres/schema

.PHONY: dev-connect dev-timer

dev-connect:
	cd apps/connect-server && air

dev-timer:
	cd apps/timer-server && air

generate-db:
	cd pkg/storage/postgres && sqlc generate

# Create a new pair of .up.sql and .down.sql files
# Usage: make migrate-new name=create_characters
migrate-new:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

# Apply all pending migrations (Bring up the database)
migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

# Rollback the last applied migration (Rollback 1 step)
migrate-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down 1

# Force a specific version (Useful if the database gets "dirty" by SQL error)
migrate-force:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose force $(version)

# Shortcut to run everything from scratch (Reset)
reset-db:
	@echo "Resetting database..."
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down -all
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up