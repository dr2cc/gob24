package app

import (
	"fmt"
	"os"

	"github.com/dr2cc/gob24/internal/bitrix/catalog"
	"github.com/dr2cc/gob24/internal/bitrix/user"
)

// Функция для печати общей справки
func printGeneralUsage() {
	fmt.Println("Утилита для работы с Битрикс24 REST API")
	fmt.Println("\nИспользование:")
	fmt.Println("  go run .\\cmd\\gob24\\main.go [команда]")
	fmt.Println("\nДоступные команды:")
	fmt.Println("  user       Добавление нового внутреннего сотрудника")
	fmt.Println("  catalog    Работа с торговым каталогом (товарами)")
	// fmt.Println("\nДля просмотра флагов конкретной команды введите:")
	// fmt.Println("  go run main.go [команда] --help")
}

func Run() error {
	// // Сейчас (17.08.2026) создаю client и ctx непосредственно в catalog.Catalog()
	// client := b24.NewClient(os.Getenv("B24_WEBHOOK_URL"))
	// ctx := context.Background()

	// Если аргументов нет вообще ИЛИ пользователь написал главные флаги помощи -h / --help
	// os.Args[0] — это путь к самому бинарнику, а os.Args[1] — первый аргумент
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		printGeneralUsage()
		os.Exit(0) // Успешный выход, так как пользователь просто попросил справку
	}

	scope := os.Args[1]

	switch scope {
	case "user":
		employeeData := os.Getenv("USER_PATH")
		fileBytes, err := os.ReadFile(employeeData)
		if err != nil {
			return fmt.Errorf("не удалось найти или прочитать файл config/user.yaml: %w", err)
		}
		fmt.Println("Запуск модуля: Пользователи")
		user.UserAdd(fileBytes)
	case "catalog":
		fmt.Println("Запуск модуля: Каталог товаров")
		catalog.Catalog()
	default:
		// Если введено что-то неизвестное, выводим ошибку и общую справку
		fmt.Printf("Ошибка: неизвестная команда '%s'\n\n", scope)
		printGeneralUsage()
		os.Exit(1) // Выход с кодом ошибки
	}
	return nil
}
