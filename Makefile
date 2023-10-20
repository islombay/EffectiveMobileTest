DB_USERNAME :=
DB_PASSWORD :=
DB_HOST :=
DB_PORT :=
DB_NAME :=

MIGRATION_DIR = ./pkg/database/migration

migrate-up:
	migrate -path $(MIGRATION_DIR) -database "postgresql:$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up
