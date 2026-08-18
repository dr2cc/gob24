package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/dr2cc/gob24/internal/bitrix/catalog"
	"github.com/dr2cc/gob24/internal/bitrix/user"
)

func Run() error {
	// // Сейчас (17.08.2026) создаю client и ctx непосредственно в catalog.Catalog()
	// client := b24.NewClient(os.Getenv("B24_WEBHOOK_URL"))
	// ctx := context.Background()

	// catalog.Catalog()
	// user.UserAdd()

	// os.Args[0] — это путь к самому бинарнику, а os.Args[1] — первый аргумент
	if len(os.Args) < 2 {
		// Создаем и возвращаем чистую ошибку
		return errors.New("укажите модуль (scope) для запуска: user|catalog")
	}

	scope := os.Args[1]

	switch scope {
	case "user":
		fmt.Println("Запуск модуля: Пользователи")
		user.UserAdd() // Ваша функция
	case "catalog":
		fmt.Println("Запуск модуля: Каталог товаров")
		catalog.Catalog() // Новая функция для каталога
	default:
		// Если скоуп не подошел, тоже возвращаем ошибку с динамическим текстом
		return fmt.Errorf("неизвестный скоуп: %s", scope)
	}
	return nil
}
