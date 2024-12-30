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
	go run -C ../cmd/migrations migrations.go -command rename -base ../../internal/infrastructures/drivers/postgres/migrations/add_items_table

mup:
	./manage_local_db_container.sh
	go run -C ../cmd/migrations migrations.go -command up \
	-db "postgres://some-postgres:mysecretpassword@localhost:5432/some-postgres?sslmode=disable" \
	-path "file://../../internal/infrastructures/drivers/postgres/migrations"
	
mdown:
	./manage_local_db_container.sh
	go run -C ../cmd/migrations migrations.go -command down -steps 1 \
	-db "postgres://some-postgres:mysecretpassword@localhost:5432/some-postgres?sslmode=disable" \
	-path "file://../../internal/infrastructures/drivers/postgres/migrations"

# to define a version of the schema for the up to start at (if force 0 -> when up: skip 0 and start at 1) or to make a version clean
mforce:
	./manage_local_db_container.sh
	go run -C ../cmd/migrations migrations.go -command force -version 1 \
	-db "postgres://some-postgres:mysecretpassword@localhost:5432/some-postgres?sslmode=disable" \
	-path "file://../../internal/infrastructures/drivers/postgres/migrations"

bmigrations:
	go build -C ../cmd/migrations -o migrations
	mv ../cmd/migrations/migrations ../build/migrations

brelease:
	go build -C ../cmd/inventory -ldflags "\
		-X main.buildMode=release \
		-X main.dbHost=localhost \
		-X main.dbUser=some-postgres \
		-X main.dbPassword=mysecretpassword \
		-X main.dbName=some-postgres \
		-X main.dbPort=5432 \
		-X main.dbSSLMode=disable" \
		-o inv-meul-app & \
	mv ../cmd/inventory/inv-meul-app ../build/inv-meul-app

bdebug:
	go build -C ../cmd/inventory -ldflags "\
		-X main.buildMode=debug \
		-X main.dbHost=localhost \
		-X main.dbUser=some-postgres \
		-X main.dbPassword=mysecretpassword \
		-X main.dbName=some-postgres \
		-X main.dbPort=5432 \
		-X main.dbSSLMode=disable" \
		-o inv-meul-app & \
	mv ../cmd/inventory/inv-meul-app ../build/inv-meul-app
