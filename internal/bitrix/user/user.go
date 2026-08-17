package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	keyboard "github.com/eiannone/keyboard"
)

// Структура запроса с учетом специфики Битрикс24
type BitrixUserFields struct {
	Email        string `json:"EMAIL"`
	Name         string `json:"NAME"`
	LastName     string `json:"LAST_NAME"`
	WorkPosition string `json:"WORK_POSITION,omitempty"` // Должность
	UFDepartment []int  `json:"UF_DEPARTMENT"`           // Массив ID отделов (например, [1])
}

// Структура для разбора ЛЮБОГО ответа от Битрикс24
type BitrixResponse struct {
	Result           int    `json:"result"`            // Сюда запишется ID, если всё ок
	ErrorType        string `json:"error"`             // Код ошибки (например, ERROR_USER_EMAIL_ALREADY_EXISTS)
	ErrorDescription string `json:"error_description"` // Понятное описание ошибки
}

func UserAdd() {
	method := "user.add"
	webhook := os.Getenv("B24_WEBHOOK_URL")
	// Если в конце нет слэша, добавляем его сами
	if webhook[len(webhook)-1] != '/' {
		webhook += "/"
	}
	url := webhook + method

	// TODO: вынести из кода
	// Заполняем данные сотрудника
	employee := BitrixUserFields{
		Email:        "tmail99@list.ru",
		Name:         "Петр",
		LastName:     "Петров",
		WorkPosition: "Менеджер по продажам",
		UFDepartment: []int{1},
	}

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
