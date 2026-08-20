package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	keyboard "github.com/eiannone/keyboard"
	"go.yaml.in/yaml/v2"
)

// Функция для чтения данных из YAML файла
func LoadUserFromYAML(fileBytes []byte) (BitrixUserFields, error) {
	var employee BitrixUserFields

	// Декодируем YAML в структуру
	err := yaml.Unmarshal(fileBytes, &employee)
	if err != nil {
		return employee, fmt.Errorf("ошибка разбора YAML: %w", err)
	}

	// 3. Валидация обязательного поля
	if employee.Email == "" {
		return employee, fmt.Errorf("в файле конфигурации не указан обязательный email")
	}

	return employee, nil
}

func UserAdd(fileBytes []byte, webhook string) {
	method := "user.add"

	employee, err := LoadUserFromYAML(fileBytes)
	if err != nil {
		fmt.Printf("Ошибка: не удалось LoadUserFromYAML: %v\n", err)
		return
	}

	// webhook := os.Getenv("B24_WEBHOOK_URL")

	// Если в конце нет слэша, добавляем его сами
	if webhook[len(webhook)-1] != '/' {
		webhook += "/"
	}
	url := webhook + method

	// TODO: выделить в отдельную функцию
	// === БЛОК ПРОВЕРКИ И ПОДТВЕРЖДЕНИЯ ===
	fmt.Println("==================================================")
	fmt.Printf("ПОДГОТОВКА К ОТПРАВКЕ ЗАПРОСА\n")
	fmt.Printf("URL назначения: %s\n\n", url)
	fmt.Println("Данные нового сотрудника:")
	fmt.Printf("  - Email:     %s\n", employee.Email)
	fmt.Printf("  - Имя:       %s\n", employee.Name)
	fmt.Printf("  - Фамилия:   %s\n", employee.LastName)
	fmt.Printf("  - Должность: %s\n", employee.WorkPosition)
	fmt.Printf("  - ID отдела: %v\n", employee.UFDepartment)
	fmt.Println("==================================================")

	fmt.Print("Нажмите [Y] для отправки или любую другую клавишу для отмены...")

	// Открываем доступ к клавиатуре
	if err := keyboard.Open(); err != nil {
		fmt.Printf("\nОшибка инициализации клавиатуры: %v\n", err)
		return
	}
	// Обязательно закрываем доступ в конце работы блока
	defer keyboard.Close()

	// Читаем нажатие одной клавиши
	char, key, err := keyboard.GetKey()
	if err != nil {
		fmt.Printf("\nОшибка чтения клавиши: %v\n", err)
		return
	}

	// Переводим строку на новую позицию, так как Enter не нажимался
	fmt.Println()

	// Проверяем: нажата ли клавиша 'y', 'Y' или специальная клавиша Enter (на всякий случай)
	if char != 'y' && char != 'Y' && key != keyboard.KeyEnter {
		fmt.Println("Выполнение отменено пользователем. Запрос не отправлен.\n\nPress any key...")
		return // Прерываем функцию
	}
	// ====================================

	jsonData, err := json.Marshal(employee)
	if err != nil {
		fmt.Printf("Ошибка JSON: %v\n", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Ошибка запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ сервера в любом случае (и при 200, и при 400)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Не удалось прочитать ответ сервера: %v\n", err)
		return
	}

	// Декодируем полученный JSON в нашу структуру
	var b24Resp BitrixResponse
	if err := json.Unmarshal(body, &b24Resp); err != nil {
		fmt.Printf("Ошибка разбора JSON ответа: %v\n", err)
		return
	}

	// Проверяем статус ответа
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Сотрудник успешно добавлен! Его ID в Битрикс24: %d\n\nPress any key...", b24Resp.Result)
	} else {
		// Если код не 200 (например, 400), выводим красивую ошибку из JSON
		fmt.Printf("Битрикс24 отклонил запрос (Код %d).\n", resp.StatusCode)
		fmt.Printf("Причина: [%s] %s\n\nPress any key...", b24Resp.ErrorType, b24Resp.ErrorDescription)
	}
}
