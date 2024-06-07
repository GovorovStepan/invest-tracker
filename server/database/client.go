package database

import (
	"log"

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
	log.Println("DB migration completed!")
}
