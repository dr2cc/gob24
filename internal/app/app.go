package app

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dr2cc/gob24/internal/bitrix/catalog"
	"github.com/dr2cc/gob24/internal/bitrix/department"
	"github.com/dr2cc/gob24/internal/bitrix/user"
	"github.com/dr2cc/gob24/internal/config"
	"github.com/dr2cc/gob24/internal/lib/logger/sl"
)

// Функция для печати общей справки
func printGeneralUsage() {
	fmt.Println("Утилита для работы с Битрикс24 REST API")
	fmt.Println("\nИспользование:")
	fmt.Println("  go run .\\cmd\\gob24\\main.go [команда]")
	fmt.Println("\nДоступные команды:")
	fmt.Println("  depts      Получение отделов в структуре организации")
	fmt.Println("  user       Добавление нового внутреннего сотрудника")
	fmt.Println("  catalog    Работа с торговым каталогом (товарами)")
	// fmt.Println("\nДля просмотра флагов конкретной команды введите:")
	// fmt.Println("  go run main.go [команда] --help")
}

func Run(cfg *config.Config) error {
	log := sl.SetupLogger(cfg.Env)
	slog.SetDefault(log)
	log.Info("starting application", slog.String("env", cfg.Env))

	// // Сейчас (17.08.2026) создаю client и ctx непосредственно в catalog.Catalog()
	// client := b24.NewClient(webhookURL)
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
		employeeData := cfg.UserPath
		fileBytes, err := os.ReadFile(employeeData)
		if err != nil {
			log.Error("не удалось найти или прочитать файл config/user.yaml:", "err", err)
			return err
		}
		fmt.Println("Запуск модуля: Пользователи")
		user.UserAdd(fileBytes, cfg.WebhookURL)
	case "catalog":
		fmt.Println("Запуск модуля: Каталог товаров")
		catalog.Catalog(cfg.WebhookURL)
	case "depts":
		// Новая команда для просмотра структуры
		if err := department.DepartmentList(cfg.WebhookURL); err != nil {
			// fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			log.Error("Ошибка:", "err", err)
		}
	default:
		// // Если введено что-то неизвестное, выводим ошибку и общую справку
		// fmt.Printf("Ошибка: неизвестная команда '%s'\n\n", scope)
		log.Error("Ошибка: неизвестная команда", "err", scope)
		printGeneralUsage()
		os.Exit(1) // Выход с кодом ошибки
	}
	return nil
}
