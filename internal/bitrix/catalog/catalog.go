package catalog

import (
	"context"
	"encoding/json"
	"fmt"

	b24 "github.com/bitrix24/b24gosdk"
	"github.com/dr2cc/gob24/internal/bitrix"
)

func Catalog(webhookURL string) error {
	client := b24.NewClient(webhookURL)
	ctx := context.Background()

	fmt.Println("⏳ Запрашиваем список торговых каталогов...")

	// Передаем в дженерик BitrixList с нашей моделью товара внутри
	data, err := bitrix.FetchList[bitrix.BitrixList[CatalogItem]](ctx, client, "catalog.catalog.list", b24.Params{
		"select": []string{"id", "iblockId", "name"},
	})

	if err != nil {
		fmt.Printf("❌ Ошибка вызова API: %v\n", err)
		return nil
	}

	// Выводим информацию на экран
	fmt.Println("\n✅ Доступные каталоги на вашем портале:")
	for _, cat := range data.Catalogs {
		fmt.Printf("----------------------------------------\n")
		fmt.Printf("📦 Название:  %s\n", cat.Name)
		fmt.Printf("🆔 iblockId:  %d  👈 ИСПОЛЬЗУЙТЕ ЭТО ЧИСЛО\n", cat.IblockID)
		fmt.Printf("🔹 Catalog ID: %d\n", cat.ID)
	}
	fmt.Printf("----------------------------------------\n")

	// ВЫНЕСТИ В ОТДЕЛЬНУЮ ФУНКЦИЮ

	// ❗❗ Готовый код со страницы
	// https://apidocs.bitrix24.ru/api-reference/catalog/product/catalog-product-add.html?tabs=defaultTabsGroup-o39srafm_go

	// client и ctx уже созданы — см. раздел «SDK для Go»
	res, err := client.Core().Call(ctx, "catalog.product.add", b24.Params{
		"fields": b24.Params{
			"iblockId":       24,
			"name":           "Товар - 17.08.2026",
			"active":         "Y",
			"barcodeMulti":   "Y",
			"canBuyZero":     "Y",
			"code":           "Tovar",
			"createdBy":      1,
			"dateActiveFrom": "2024-05-28T10:00:00",
			"dateActiveTo":   "2024-05-29T10:00:00",
			"dateCreate":     "2024-05-27T10:00:00",
			"detailPicture": b24.Params{
				"fileData": []string{"detailPicture.png", "iVBORw0KGgoAAAANSUhEUgAAAMgAAADIBAMAAABfdrOtAAAAG1BMVEX37ff/­///58fn9+v3+/P779vv8+Pz47/j68/oDfe+3AAAACXBIWXMAAA7EAAAOxAGV­Kw4bAAABrUlEQVR4nO3UT0/CMBjH8ccx2I56IFynkHg1SgxHHCocSfQFGKP3­+e++xL1wn7bPUCAeKF5Mvp+EluX3ZN3ariIAAAAAAAAAAAAAAAAA/q2TwrXZ­ib94LTbj5GdgVbtKxhdXS+2uL270ajQbL9fz4WzcXwVWtbNeIdmt3qSQtwdJ­Ssku1/NHkfdVEKriHFey0G4haS3+ty4ZtEGoipMW+VS7T2m0zc+28tICq4rT­qXtuJV7kWdvsUJtuoc1Hm08ssKo4B1Wn1i6tJu5qrj9dA8lWEzOQEFhV3CCN­Tph2naJ0V+eu0SV+ry3WWQqBVcUNsgiP16ndS4SnzuffL5LWEgKrihqje7Y9­iDTN6mZ38geDNNX2dEm338b5XPafrmRuj/dj4fULfGoXeFTJ/guvayybW1i3­Vl7aM7h+3y2c+y07FfeZjaT9GHVrNYXPG/fkIbCqCPf+9d1WKiWtJSyP21r+­FaTrZ8+CULW7XliCUe0PyIUdkD29qQzdv7A0FoSq3R0fqaU78d0hPtw86hMX­99vAqqJlp757/W3vhMCqAAAAAAAAAAAAAAAAAPxbX82/SILlk9xfAAAAAElF­TkSuQmCCiVBORw0KGgoAAAANSUhEUgAAAMgAAADIBAMAAABfdrOtAAAAG1BM­VEX37ff////58fn9+v3+/P779vv8+Pz47/j68/oDfe+3AAAACXBIWXMAAA7E­AAAOxAGVKw4bAAABrUlEQVR4nO3UT0/CMBjH8ccx2I56IFynkHg1SgxHHCoc­SfQFGKP3+e++xL1wn7bPUCAeKF5Mvp+EluX3ZN3ariIAAAAAAAAAAAAAAAAA­/q2TwrXZib94LTbj5GdgVbtKxhdXS+2uL270ajQbL9fz4WzcXwVWtbNeIdmt­3qSQtwdJSsku1/NHkfdVEKriHFey0G4haS3+ty4ZtEGoipMW+VS7T2m0zc+2­8tICq4rTqXtuJV7kWdvsUJtuoc1Hm08ssKo4B1Wn1i6tJu5qrj9dA8lWEzOQ­EFhV3CCNTph2naJ0V+eu0SV+ry3WWQqBVcUNsgiP16ndS4SnzuffL5LWEgKr­ihqje7Y9iDTN6mZ38geDNNX2dEm338b5XPafrmRuj/dj4fULfGoXeFTJ/guv­ayybW1i3Vl7aM7h+3y2c+y07FfeZjaT9GHVrNYXPG/fkIbCqCPf+9d1WKiWt­JSyP21r+FaTrZ8+CULW7XliCUe0PyIUdkD29qQzdv7A0FoSq3R0fqaU78d0h­Ptw86hMX99vAqqJlp757/W3vhMCqAAAAAAAAAAAAAAAAAPxbX82/SILlk9xf­AAAAAElFTkSuQmCC"},
			},
			"detailText":      "",
			"detailTextType":  "text",
			"height":          100,
			"iblockSectionId": 47,
			"length":          100,
			"measure":         5,
			"modifiedBy":      1,
			"previewPicture": b24.Params{
				"fileData": []string{"previewPicture.png", "iVBORw0KGgoAAAANSUhEUgAAAMgAAADIBAMAAABfdrOtAAAAG1BMVEX37ff/­///58fn9+v3+/P779vv8+Pz47/j68/oDfe+3AAAACXBIWXMAAA7EAAAOxAGV­Kw4bAAABrUlEQVR4nO3UT0/CMBjH8ccx2I56IFynkHg1SgxHHCocSfQFGKP3­+e++xL1wn7bPUCAeKF5Mvp+EluX3ZN3ariIAAAAAAAAAAAAAAAAA/q2TwrXZ­ib94LTbj5GdgVbtKxhdXS+2uL270ajQbL9fz4WzcXwVWtbNeIdmt3qSQtwdJ­Ssku1/NHkfdVEKriHFey0G4haS3+ty4ZtEGoipMW+VS7T2m0zc+28tICq4rT­qXtuJV7kWdvsUJtuoc1Hm08ssKo4B1Wn1i6tJu5qrj9dA8lWEzOQEFhV3CCN­Tph2naJ0V+eu0SV+ry3WWQqBVcUNsgiP16ndS4SnzuffL5LWEgKrihqje7Y9­iDTN6mZ38geDNNX2dEm338b5XPafrmRuj/dj4fULfGoXeFTJ/guvayybW1i3­Vl7aM7h+3y2c+y07FfeZjaT9GHVrNYXPG/fkIbCqCPf+9d1WKiWtJSyP21r+­FaTrZ8+CULW7XliCUe0PyIUdkD29qQzdv7A0FoSq3R0fqaU78d0hPtw86hMX­99vAqqJlp757/W3vhMCqAAAAAAAAAAAAAAAAAPxbX82/SILlk9xfAAAAAElF­TkSuQmCCiVBORw0KGgoAAAANSUhEUgAAAMgAAADIBAMAAABfdrOtAAAAG1BM­VEX37ff////58fn9+v3+/P779vv8+Pz47/j68/oDfe+3AAAACXBIWXMAAA7E­AAAOxAGVKw4bAAABrUlEQVR4nO3UT0/CMBjH8ccx2I56IFynkHg1SgxHHCoc­SfQFGKP3+e++xL1wn7bPUCAeKF5Mvp+EluX3ZN3ariIAAAAAAAAAAAAAAAAA­/q2TwrXZib94LTbj5GdgVbtKxhdXS+2uL270ajQbL9fz4WzcXwVWtbNeIdmt­3qSQtwdJSsku1/NHkfdVEKriHFey0G4haS3+ty4ZtEGoipMW+VS7T2m0zc+2­8tICq4rTqXtuJV7kWdvsUJtuoc1Hm08ssKo4B1Wn1i6tJu5qrj9dA8lWEzOQ­EFhV3CCNTph2naJ0V+eu0SV+ry3WWQqBVcUNsgiP16ndS4SnzuffL5LWEgKr­ihqje7Y9iDTN6mZ38geDNNX2dEm338b5XPafrmRuj/dj4fULfGoXeFTJ/guv­ayybW1i3Vl7aM7h+3y2c+y07FfeZjaT9GHVrNYXPG/fkIbCqCPf+9d1WKiWt­JSyP21r+FaTrZ8+CULW7XliCUe0PyIUdkD29qQzdv7A0FoSq3R0fqaU78d0h­Ptw86hMX99vAqqJlp757/W3vhMCqAAAAAAAAAAAAAAAAAPxbX82/SILlk9xf­AAAAAElFTkSuQmCC"},
			},
			"previewText":        "",
			"previewTextType":    "text",
			"purchasingCurrency": "RUB",
			"purchasingPrice":    1000,
			"quantity":           10,
			"quantityReserved":   1,
			"quantityTrace":      "Y",
			"recurSchemeLength":  1,
			"recurSchemeType":    "D",
			"sort":               100,
			"subscribe":          "Y",
			"trialPriceId":       175,
			// "vatId":              1,
			// "vatIncluded":        "Y",
			// "weight":             100,
			// "width":              100,
			// "withoutOrder":       "Y",
			// "xmlId":              "",
			// "property258":        "test",
			// "property259":        []string{"test1", "test2"},
		},
	})
	if err != nil {
		return fmt.Errorf("catalog.product.add: %w", err)
	}

	// Метод заворачивает ответ в объект с ключом "element".
	raw, ok := b24.Unwrap(res.Result, "element")
	if !ok {
		return fmt.Errorf("в ответе нет ключа element")
	}

	var item struct {
		Active     string `json:"active"`
		Available  string `json:"available"`
		Bundle     string `json:"bundle"`
		CanBuyZero string `json:"canBuyZero"`
		Code       string `json:"code"`
		CreatedBy  int    `json:"createdBy"`
	}
	if err := json.Unmarshal(raw, &item); err != nil {
		return fmt.Errorf("разбор ответа: %w", err)
	}
	fmt.Println(item.Active, item.Available)

	return nil
}
