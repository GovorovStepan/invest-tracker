package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/models"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to DB")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Settings{})
	Instance.AutoMigrate(&models.Portfolio{})
	Instance.AutoMigrate(&models.Position{})
	Instance.AutoMigrate(&models.Transaction{})
	log.Println("DB migration completed!")
}

func FormatConnectionString() string {
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable "
	}
	timezone := os.Getenv("DB_TIMEZONE")
	if timezone == "" {
		timezone = "UTC"
	}
	dbHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		log.Fatal("DB_HOST is not set")
	}
	dbUser, exists := os.LookupEnv("DB_USERNAME")
	if !exists {
		log.Fatal("DB_USERNAME is not set")
	}
	dbPass, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Fatal("DB_PASSWORD is not set")
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		log.Fatal("DB_NAME is not set")
	}

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, port, sslmode, timezone)

	return connectionString
}
