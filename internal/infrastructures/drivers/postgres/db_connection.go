package infrastructures_drivers_postgres

import (
	"log"
	"net/url"
	"os"
	"strings"

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

func DSNToConnectionString(dsn string) string {
	var connectionStringBuilder strings.Builder
	dsnParts := strings.Fields(dsn)
	dsnMap := make(map[string]string)

	for _, part := range dsnParts {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 2 {
			dsnMap[kv[0]] = kv[1]
		}
	}

	if user, ok := dsnMap["user"]; ok {
		connectionStringBuilder.WriteString("postgres://" + user)
	}
	if password, ok := dsnMap["password"]; ok {
		connectionStringBuilder.WriteString(":" + password + "@")
	}
	if host, ok := dsnMap["host"]; ok {
		connectionStringBuilder.WriteString(host)
	}
	if port, ok := dsnMap["port"]; ok {
		connectionStringBuilder.WriteString(":" + port)
	}
	if dbname, ok := dsnMap["dbname"]; ok {
		connectionStringBuilder.WriteString("/" + dbname)
	}

	options := url.Values{}

	for key, value := range dsnMap {

		if key != "user" && key != "password" && key != "host" && key != "port" && key != "dbname" {
			options.Add(key, value)
		}
	}

	if encodedOptions := options.Encode(); encodedOptions != "" {
		connectionStringBuilder.WriteString("?" + encodedOptions)
	}
	connectionString := connectionStringBuilder.String()
	return connectionString
}
