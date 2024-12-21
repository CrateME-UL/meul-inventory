i:
	go mod tidy

r:
	go run ../cmd/inventory/inventory.go

brelease:
	go build -C ../cmd/inventory -ldflags "-X main.buildMode=release" -o inv-meul-app & \
	mv ../cmd/inventory/inv-meul-app ../build/inv-meul-app

bdebug:
	go build -C ../cmd/inventory -ldflags "-X main.buildMode=debug" -o inv-meul-app & \
	mv ../cmd/inventory/inv-meul-app ../build/inv-meul-app
