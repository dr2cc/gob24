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

	// Run
	if err := app.Run(); err != nil {
		// Более сложная обработка ошибки
		// для дальнейшего внедрения graceful shutdown
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
