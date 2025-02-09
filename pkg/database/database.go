package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB


type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Logger   logger.Interface
}


func Connect(databaseConfig DatabaseConfig) error {
	var err error

	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Name,
	)

	
	if databaseConfig.Logger == nil {
		databaseConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: databaseConfig.Logger,
	})

	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return err
	}

	
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to get sql.DB from gorm DB instance: %v", err)
		return err
	}

	
	sqlDB.SetMaxIdleConns(10)   
	sqlDB.SetMaxOpenConns(100) 
	sqlDB.SetConnMaxLifetime(60) 

	return nil
}
