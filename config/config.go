package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConfig() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDb != nil {
		panic("Error connecting to database")
	}
	fmt.Printf("Connected to database %s\n", dbName)

	return db
}
