// Package rest provides the http rest api
package rest

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type RestConfig struct {
	BuildMode    string
	TrustedProxy []string
	Port         string
}

// RouteRegisterFunc is a function type for registering routes
type RouteRegisterFunc func(*gin.Engine)

func DefaultRestServer(restConfig RestConfig, routeRegisterFuncs []RouteRegisterFunc) (*gin.Engine, error) {
	if restConfig.BuildMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	err := r.SetTrustedProxies(restConfig.TrustedProxy)
	if err != nil {
		fmt.Printf("failed to set trusted proxies: %s\n", err)
		os.Exit(2)
	}

	// Execute all registration functions to register routes
	for _, registeredFunction := range routeRegisterFuncs {
		registeredFunction(r)
	}

	//setup the static directory from the web directory as root
	r.Static("/static", "./static")

	return r, nil
}