package infrastructures

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
