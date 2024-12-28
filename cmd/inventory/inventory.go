package main

import (
	"fmt"
	infrastructures_drivers_postgres "meul/inventory/internal/infrastructures/drivers/postgres"
	rest "meul/inventory/internal/interfaces/rest"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	buildMode    string
	port         string = ":3000"
	trustedProxy        = []string{"127.0.0.1", "::1"}
	dbHost       string
	dbUser       string
	dbPassword   string
	dbName       string
	dbPort       string
	dbSSLMode    string
	dbTimeZone   string = "America/Toronto"
)

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
		fmt.Printf("failed to set trusted proxies: %s\n", err)
		os.Exit(2)
	}

	rest.RegisterRoutes(r)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone,
	)
	dbConfig := infrastructures_drivers_postgres.DbConfig{
		DSN: dsn,
	}
	e, err := InitializeEvent(dbConfig)
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()

	r.Run(port)
}
