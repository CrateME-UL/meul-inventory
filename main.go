package main

import (
	routes "monequillibreul/inventory/res"

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

	routes.RegisterRoutes(r)

	r.Run(port)
}
