package config

import (
	"Gofinal1/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=expense_management port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// ทำการ migrate ตาราง
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Expense{})

	DB = database
}
