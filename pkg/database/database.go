package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Config is the configuration for the database
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Logger   logger.Interface
}

// Connect establishes a connection to the database with dynamic config
func Connect(config Config) error {
	var err error

	// Build the Data Source Name (DSN) string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Name,
	)

	// Set the logger if provided, else use the default logger
	if config.Logger == nil {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	// Open the database connection using the config values
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: config.Logger,
	})

	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return err
	}

	// Optional: Set other database connection pool options
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to get sql.DB from gorm DB instance: %v", err)
		return err
	}

	// Set connection pool options (optional)
	sqlDB.SetMaxIdleConns(10)   // Set max idle connections
	sqlDB.SetMaxOpenConns(100)  // Set max open connections
	sqlDB.SetConnMaxLifetime(60) // Set max connection lifetime (seconds)

	return nil
}
