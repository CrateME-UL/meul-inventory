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
	go run -C ../cmd/migrations migrations.go -command rename -base ../../internal/infrastructures/migrations/add_items_table

m:
	go run migrations.go -command up -db "postgres://some-postgres:mysecretpassword@localhost:5432/some-postgres?sslmode=disable" \
	-path "file://../internal/infrastructures/migrations"


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
