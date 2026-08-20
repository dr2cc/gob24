package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dr2cc/gob24/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из файла .env в корне
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	// TODO: вынести в config
	// Инициализируем клиент с проверкой переменной окружения
	webhookURL := os.Getenv("B24_WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("❌ Ошибка: переменная окружения B24_WEBHOOK_URL не задана")
	}

	// Run
	if err := app.Run(webhookURL); err != nil {
		// Более сложная обработка ошибки
		// для дальнейшего внедрения graceful shutdown
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
