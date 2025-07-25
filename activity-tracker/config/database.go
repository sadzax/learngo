package config

import (
	"fmt"
	"log"
	"os"

	"activity-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // symlynk alike?

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), // DB_PASSWORD другой а в dsn у меня  DSN is: host=localhost user=sadzax password=pard dbname=activity_tracker port=5432 sslmode=disable, type is string
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	fmt.Println("DSN formed \n")
	fmt.Printf("DSN is: %v, type is %T\n", dsn, dsn)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Автомиграция (HANDLE ERROR!)
	database.AutoMigrate(&models.User{}, &models.Activity{}, &models.ActivityLog{})

	DB = database
}
