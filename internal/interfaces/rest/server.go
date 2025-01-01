// Package rest provides ...
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

func DefaultRestServer(restConfig RestConfig) (*gin.Engine, error) {
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

	RegisterRoutes(r)

	return r, nil
}
