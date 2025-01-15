i:
	go mod tidy

r:
	make i
	cd ../cmd/inventory && wire
	./manage_local_db_container.sh
	go run -ldflags "\
		-X main.buildMode=debug \
		-X main.dbHost=localhost \
		-X main.dbUser=some-postgres \
		-X main.dbPassword=mysecretpassword \
		-X main.dbName=some-postgres \
		-X main.dbPort=5432 \
		-X main.dbSSLMode=disable" \
		../cmd/inventory/inventory.go ../cmd/inventory/wire_gen.go

mrename:
	cd ../cmd/migrations && wire
	go run -C ../cmd/migrations migrations.go -command rename \
	-base ../../internal/infrastructures/drivers/postgres/migrations/sql/add_items_table

mup:
	cd ../cmd/migrations && wire
	./manage_local_db_container.sh
	go run -ldflags "\
		-X main.buildMode=debug \
		-X main.dbHost=localhost \
		-X main.dbUser=some-postgres \
		-X main.dbPassword=mysecretpassword \
		-X main.dbName=some-postgres \
		-X main.dbPort=5432 \
		-X main.dbSSLMode=disable" \
	../cmd/migrations/migrations.go ../cmd/migrations/wire_gen.go \
	-command up \
	-path "file://../internal/infrastructures/drivers/postgres/migrations/sql"
	
mdown:
	cd ../cmd/migrations && wire
	./manage_local_db_container.sh
	go run -ldflags "\
		-X main.buildMode=debug \
		-X main.dbHost=localhost \
		-X main.dbUser=some-postgres \
		-X main.dbPassword=mysecretpassword \
		-X main.dbName=some-postgres \
		-X main.dbPort=5432 \
		-X main.dbSSLMode=disable" \
	../cmd/migrations/migrations.go ../cmd/migrations/wire_gen.go \
	-command down \
	-steps 1 \
	-path "file://../internal/infrastructures/drivers/postgres/migrations/sql"

# to define a version of the schema for the up to start at (if force 0 -> when up: skip 0 and start at 1) or to make a version clean
mforce:
	cd ../cmd/migrations && wire
	./manage_local_db_container.sh
	go run -C ../cmd/migrations migrations.go wire_gen.go -command force -version 1 \
	-db "postgres://some-postgres:mysecretpassword@localhost:5432/some-postgres?sslmode=disable" \
	-path "file://../../internal/infrastructures/drivers/postgres/migrations/sql"

bmigrations:
	go build -C ../cmd/migrations -o migrations
	mv ../cmd/migrations/migrations ../build/migrations

brelease:
	make i
	cd ../cmd/inventory && wire
	go build -C ../cmd/inventory -ldflags "\
		-X main.buildMode=release \
		-X main.dbHost=localhost \
		-X main.dbUser=some-postgres \
		-X main.dbPassword=mysecretpassword \
		-X main.dbName=some-postgres \
		-X main.dbPort=5432 \
		-X main.dbSSLMode=disable" \
		-o inv-meul-app inventory.go wire_gen.go && mv ../cmd/inventory/inv-meul-app ../build/inv-meul-app

bdebug:
	make i
	cd ../cmd/inventory && wire
	go build -C ../cmd/inventory -ldflags "\
		-X main.buildMode=debug \
		-X main.dbHost=localhost \
		-X main.dbUser=some-postgres \
		-X main.dbPassword=mysecretpassword \
		-X main.dbName=some-postgres \
		-X main.dbPort=5432 \
		-X main.dbSSLMode=disable" \
		-o inv-meul-app inventory.go wire_gen.go && mv ../cmd/inventory/inv-meul-app ../build/inv-meul-app

# live server to enable hot reload on the html templates
air:
	./manage_local_db_container.sh
	cd .. && air

static:
	npm --prefix ../web install
	npm --prefix ../web run build
	cp ../web/node_modules/htmx.org/dist/htmx.min.js ../web/static/js/htmx.min.js
	make tw

tw:
	cd ../web && npx tailwindcss -i ./src/input.css -o ./static/css/output.css --watch

