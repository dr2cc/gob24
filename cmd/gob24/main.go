package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dr2cc/gob24/internal/app"
	"github.com/dr2cc/gob24/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	if err := app.Run(cfg); err != nil {
		// Более сложная обработка ошибки
		// для дальнейшего внедрения graceful shutdown
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
