i:
	go mod tidy

r:
	go run main.go

brelease:
	go build -ldflags "-X main.buildMode=release" -o inv-meul-app

bdebug:
	go build -ldflags "-X main.buildMode=debug" -o inv-meul-app
