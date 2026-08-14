package bitrix

import (
	"context"
	"encoding/json"
	"fmt"

	b24 "github.com/bitrix24/b24gosdk"
)

// Обертка списков будет принимать срез элементов
type BitrixList[T any] struct {
	Catalogs []T `json:"catalogs"` // Для catalog.catalog.list
	// Сюда же потом можно добавить теги json для других методов, если они отличаются
}

// Функция FetchList теперь распаковывает JSON напрямую в ожидаемый тип T
func FetchList[T any](ctx context.Context, client *b24.Client, method string, params b24.Params) (T, error) {
	var result T

	// 1. Делаем вызов в Битрикс24
	res, err := client.Core().Call(ctx, method, params)
	if err != nil {
		return result, fmt.Errorf("ошибка вызова метода %s: %w", method, err)
	}

	// 2. Распаковываем res.Result НАПРЯМУЮ в result (без ListResponse)
	if err := json.Unmarshal(res.Result, &result); err != nil {
		return result, fmt.Errorf("ошибка разбора JSON для метода %s: %w", method, err)
	}

	// 3. Возвращаем заполненный результат
	return result, nil
}
