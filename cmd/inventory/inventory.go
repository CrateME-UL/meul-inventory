package main

import (
	"fmt"
	rest "meul/inventory/internal/interfaces/rest"
	"os"

	"github.com/gin-gonic/gin"
)

var buildMode string
var port string = ":3000"
var trustedProxy []string = []string{"127.0.0.1", "::1"}

func main() {

	if buildMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	err := r.SetTrustedProxies(trustedProxy)
	if err != nil {
		panic(err)
	}

	rest.RegisterRoutes(r)

	e, err := InitializeEvent("hi there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()

	r.Run(port)
}
