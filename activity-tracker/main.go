package main

import (
	"activity-tracker/config"
	"log"

	"activity-tracker/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден")
	}

	// Подключаемся к базе данных
	config.ConnectDatabase()

	// Настраиваем маршруты
	r := routes.SetupRoutes()

	// Запускаем сервер
	r.Run(":8080") // HANDLE ERROR!!!!!
}
