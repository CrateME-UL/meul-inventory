# Variables
GO_BUILD_FLAGS = -ldflags "\
	-X main.buildMode=$(BUILD_MODE) \
	-X main.dbHost=$(DB_HOST) \
	-X main.dbUser=$(DB_USER) \
	-X main.dbPassword=$(DB_PASSWORD) \
	-X main.dbName=$(DB_NAME) \
	-X main.dbPort=$(DB_PORT) \
	-X main.dbSSLMode=$(DB_SSL_MODE)"
DB_HOST = localhost
DB_USER = some-postgres
DB_PASSWORD = mysecretpassword
DB_NAME = some-postgres
DB_PORT = 5432
DB_SSL_MODE = disable
BUILD_DIR = ../build
CMD_DIR = ../cmd
MIGRATIONS_PATH = file://../internal/infrastructures/drivers/postgres/migrations/sql

# Default target
default: i

t:
	cd $(CMD_DIR)/inventory && wire
	cd $(CMD_DIR)/migrations && wire
	go test ../...

# Install dependencies
i:
	go mod tidy

# Run the inventory service
r: i
	cd $(CMD_DIR)/inventory && wire
	./manage_local_db_container.sh
	go run $(GO_BUILD_FLAGS) $(CMD_DIR)/inventory/inventory.go $(CMD_DIR)/inventory/wire_gen.go

# Run migrations commands
mrename:
	cd $(CMD_DIR)/migrations && wire
	go run -C $(CMD_DIR)/migrations migrations.go -command rename \
	-base ../../internal/infrastructures/drivers/postgres/migrations/sql/add_items_table

mup:
	cd $(CMD_DIR)/migrations && wire
	./manage_local_db_container.sh
	go run $(GO_BUILD_FLAGS) $(CMD_DIR)/migrations/migrations.go $(CMD_DIR)/migrations/wire_gen.go \
	-command up \
	-path "$(MIGRATIONS_PATH)"

mdown:
	cd $(CMD_DIR)/migrations && wire
	./manage_local_db_container.sh
	go run $(GO_BUILD_FLAGS) $(CMD_DIR)/migrations/migrations.go $(CMD_DIR)/migrations/wire_gen.go \
	-command down \
	-steps 1 \
	-path "$(MIGRATIONS_PATH)"

mforce:
	cd $(CMD_DIR)/migrations && wire
	./manage_local_db_container.sh
	go run -C $(CMD_DIR)/migrations migrations.go wire_gen.go -command force -version 1 \
	-db "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" \
	-path "$(MIGRATIONS_PATH)"

# Build binaries
bmigrations:
	go build -C $(CMD_DIR)/migrations -o migrations
	mv $(CMD_DIR)/migrations/migrations $(BUILD_DIR)/migrations

brelease: BUILD_MODE = release
brelease: i
	cd $(CMD_DIR)/inventory && wire
	go build -C $(CMD_DIR)/inventory $(GO_BUILD_FLAGS) -o inv-meul-app inventory.go wire_gen.go
	mv $(CMD_DIR)/inventory/inv-meul-app $(BUILD_DIR)/inv-meul-app

bdebug: BUILD_MODE = debug
bdebug: i
	cd $(CMD_DIR)/inventory && wire
	go build -C $(CMD_DIR)/inventory $(GO_BUILD_FLAGS) -o inv-meul-app inventory.go wire_gen.go
	mv $(CMD_DIR)/inventory/inv-meul-app $(BUILD_DIR)/inv-meul-app

# Live server for hot reload
air:
	./manage_local_db_container.sh
	cd .. && air

# Build static files
static:
	npm --prefix ../web install
	npm --prefix ../web run build
	cp ../web/node_modules/htmx.org/dist/htmx.min.js ../web/static/js/htmx.min.js
	make tw

tw:
	cd ../web && npx tailwindcss -i ./src/input.css -o ./static/css/output.css --watch

prech:
	sudo apt update
	sudo apt install pipx
	pipx install pre-commit
	pipx ensurepath
	pre-commit install
	pre-commit autoupdate

l:
	cd $(CMD_DIR)/inventory && wire
	cd $(CMD_DIR)/migrations && wire
	pre-commit run --all-files







