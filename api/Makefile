# API Service Makefile
# Provides commands for database migrations and application management

#------------------------------------------------------
# Environment Configuration
#------------------------------------------------------

# Load environment variables from .env file if present
ifneq (,$(wildcard .env))
  include .env
  export
endif

#------------------------------------------------------
# Default Configuration Variables
#------------------------------------------------------

# Database connection string (override in .env)
DB_URL ?= postgres://postgres:secret@localhost:5432/mydb?sslmode=disable

# Path to database migration files
MIGRATIONS_DIR ?= manifest/migrations

# Detect migrate binary location
MIGRATE_BIN := $(shell which migrate)

#------------------------------------------------------
# Migration Commands
#------------------------------------------------------

# Verify migrate CLI is installed
check-migrate:
ifeq ($(MIGRATE_BIN),)
	@echo ""
	@echo "❌ Migrate CLI not found. Please install it manually:"
	@echo ""
	@echo "  ▸ macOS   : brew install golang-migrate"
	@echo "  ▸ Ubuntu  : curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz && sudo mv migrate /usr/local/bin"
	@echo "  ▸ Windows : https://github.com/golang-migrate/migrate/releases"
	@echo "  ▸ Any OS  : go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
	@echo ""
	@exit 1
else
	@echo ✅ Migrate found at: $(MIGRATE_BIN)
endif

# Create a new migration file with timestamp prefix
migrate-create:
	@read -p "Migration name: " name; \
	TIMESTAMP=$$(date +%Y%m%d%H%M%S); \
	mkdir -p $(MIGRATIONS_DIR); \
	touch $(MIGRATIONS_DIR)/$${TIMESTAMP}_$${name}.up.sql; \
	touch $(MIGRATIONS_DIR)/$${TIMESTAMP}_$${name}.down.sql; \
	echo "✅ Created migration: $${TIMESTAMP}_$${name}"

# Apply all pending migrations
migrate-up: check-migrate
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Rollback the most recent migration
migrate-down: check-migrate
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

# Display current migration version
migrate-status: check-migrate
	$(MIGRATE_BIN) -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

#------------------------------------------------------
# Application Commands
#------------------------------------------------------

# Start the application in development mode
run:
	go run main.go

# Build the application for production
build:
	go build -o bin/api main.go

# Run tests
test:
	go test -v ./...

#------------------------------------------------------
# Helper Commands
#------------------------------------------------------

# List all available commands
help:
	@echo "FastKit API Makefile Commands:"
	@echo "  migrate-create  - Create a new migration file"
	@echo "  migrate-up      - Apply all pending migrations"
	@echo "  migrate-down    - Rollback the most recent migration"
	@echo "  migrate-status  - Show current migration version"
	@echo "  run             - Start the application in development mode"
	@echo "  build           - Build the application for production"
	@echo "  test            - Run all tests"
	@echo "  help            - Display this help message"

.PHONY: check-migrate migrate-create migrate-up migrate-down migrate-status run build test help
