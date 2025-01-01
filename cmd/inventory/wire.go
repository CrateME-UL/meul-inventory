//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"meul/inventory/internal/infrastructures/drivers/postgres"
	"meul/inventory/internal/interfaces/rest"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
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
	DBClient   *gorm.DB
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

func DefaultApp(HttpServer *gin.Engine, DBClient *gorm.DB) (*App, error) {
	return &App{
		HttpServer: HttpServer,
		DBClient:   DBClient,
	}, nil
}

// WireSet is a set that includes all necessary providers
var WireSet = wire.NewSet(
	ProvideDBConfig,
	ProvideRestConfig,
	postgres.NewDatabaseConnection,
	rest.DefaultRestServer,
	DefaultApp,
)

func InitializeInventoryHandler() (*App, error) {
	wire.Build(WireSet)

	return nil, nil
}
