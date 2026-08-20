package department

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// FlexInt — кастомный "умный" тип, который умеет парсить из JSON как строки так и числа
type FlexInt int

// UnmarshalJSON — кастомный парсер для нашего типа
func (fi *FlexInt) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		*fi = 0
		return nil
	}

	// Если это строка (начинается с кавычки), убираем кавычки и конвертируем
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		if s == "" {
			*fi = 0
			return nil
		}
		val, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		*fi = FlexInt(val)
		return nil
	}

	// Если это обычное число
	var i int
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	*fi = FlexInt(i)
	return nil
}

func DepartmentList(webhook string) error {
	method := "department.get"
	// webhook := os.Getenv("B24_WEBHOOK_URL")
	if webhook == "" {
		return fmt.Errorf("переменная окружения B24_WEBHOOK_URL не задана")
	}
	if webhook[len(webhook)-1] != '/' {
		webhook += "/"
	}
	url := webhook + method

	// Отправляем пустой POST-запрос (метод get не требует обязательных полей)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return fmt.Errorf("ошибка запроса к Битрикс24: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("не удалось прочитать ответ сервера: %w", err)
	}

	var b24Resp BitrixDepartmentResponse
	if err := json.Unmarshal(body, &b24Resp); err != nil {
		return fmt.Errorf("ошибка разбора JSON: %w", err)
	}

	// Проверяем статус ответа сервера
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("битрикс24 вернул ошибку [%s]: %s", b24Resp.ErrorType, b24Resp.ErrorDescription)
	}

	// Выводим результат в консоль с анализом подчиненности
	fmt.Println("==================================================")
	fmt.Println("СПИСОК ОТДЕЛОВ БИТРИКС24:")
	fmt.Println("==================================================")

	for _, dept := range b24Resp.Result {
		isSubordinated := false
		parentID := 0

		// Проверяем поле PARENT. Если там число или строка (не false и не 0) — отдел подчиненный
		if dept.Parent != nil {
			switch v := dept.Parent.(type) {
			case string:
				if v != "" && v != "0" {
					isSubordinated = true
					fmt.Sscanf(v, "%d", &parentID)
				}
			case float64: // JSON numbers unmarshal to float64 by default
				if v > 0 {
					isSubordinated = true
					parentID = int(v)
				}
			}
		}

		// Формируем красивый вывод для пользователя
		statusText := "^ Корневой отдел"
		if isSubordinated {
			statusText = fmt.Sprintf("--> Подчиненный (Главный отдел ID: %d)", parentID)
		}

		// Приводим к int() при выводе на экран
		fmt.Printf("ID: %d | Название: %-25s | Статус: %s\n", int(dept.ID), dept.Name, statusText)
		// fmt.Printf("ID: %d | Название: %-25s | Статус: %s\n", dept.ID, dept.Name, statusText)
	}
	fmt.Println("==================================================")

	return nil
}
