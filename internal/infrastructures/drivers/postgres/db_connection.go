package infrastructures_drivers_postgres

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DbConfig holds the configuration for initializing components.
type DbConfig struct {
	DSN string
}

func NewDatabaseConnection(config DbConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDatabaseConnectionWithMigrationLogger(config DbConfig, logFile *os.File) (*gorm.DB, error) {
	// Set up the custom logger to write to a file
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // Log output to file
		logger.Config{
			SlowThreshold:             0,           // Disable slow query logging
			LogLevel:                  logger.Info, // Log all SQL queries
			IgnoreRecordNotFoundError: true,        // Ignore not found errors
			Colorful:                  false,       // Disable colors in logs
		},
	)

	// Set up the database connection with GORM
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{
		Logger: newLogger, // Use the custom logger
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
