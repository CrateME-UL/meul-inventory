package main

import (
	"log"
	"meul/inventory/web"
)

func main() {
	app, err := InitializeInventoryHandler()
	if err != nil {
		log.Fatalf("failed to initialize http server: %v", err)
	}

	app.HttpServer.HTMLRender = web.NewTemplate()

	app.HttpServer.Run(port)
}
