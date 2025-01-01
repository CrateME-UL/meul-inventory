package main

import "log"

func main() {
	app, err := InitializeInventoryHandler()
	if err != nil {
		log.Fatalf("failed to initialize http server: %v", err)
	}

	app.HttpServer.Run(port)
}
