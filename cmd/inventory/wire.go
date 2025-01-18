//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"meul/inventory/internal/infrastructures/drivers/postgres"
	"meul/inventory/internal/infrastructures/drivers/postgres/models"
	"meul/inventory/internal/interfaces/rest"
	"meul/inventory/internal/interfaces/rest/items"
	"meul/inventory/internal/interfaces/rest/ping"
	"meul/inventory/internal/interfaces/rest/root"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var (
	buildMode    string
	port         string   = ":3000"
	trustedProxy []string = []string{"127.0.0.1", "::1"}
	dbHost       string
	dbUser       string
	dbPassword   string
	dbName       string
	dbPort       string
	dbSSLMode    string
	dbTimeZone   string = "America/Toronto"
)

type App struct {
	HttpServer *gin.Engine
}

// ProvideDBConfig creates a DbConfig for PostgreSQL
func ProvideDBConfig() (postgres.DbConfig, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone,
	)
	dbConfig := postgres.DbConfig{
		DSN: dsn,
	}

	return dbConfig, nil
}

// ProvideRestConfig provides the parameters for the Gin server
func ProvideRestConfig() (rest.RestConfig, error) {
	return rest.RestConfig{
		BuildMode:    buildMode,
		TrustedProxy: trustedProxy,
		Port:         port,
	}, nil
}

// ProvideServerDependencies provides the dependencies for the server in a callback fashion
func ProvideServerDependencies(itemDAO *models.ItemDAO) []rest.RouteRegisterFunc {
	return []rest.RouteRegisterFunc{
		func(r *gin.Engine) { root.RegisterRoot(r) },
		func(r *gin.Engine) { ping.RegisterPing(r) },
		func(r *gin.Engine) { items.RegisterItems(r, itemDAO) },
	}
}

func DefaultApp(httpServer *gin.Engine) (*App, error) {
	return &App{
		HttpServer: httpServer,
	}, nil
}

// WireSet is a set that includes all necessary providers
var WireSet = wire.NewSet(
	ProvideDBConfig,
	ProvideRestConfig,
	ProvideServerDependencies,
	models.NewItemDAO,
	postgres.NewDatabaseConnection,
	rest.DefaultRestServer,
	DefaultApp,
)

func InitializeInventoryHandler() (*App, error) {
	wire.Build(WireSet)

	return nil, nil
}
