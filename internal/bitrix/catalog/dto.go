package catalog

// Структура каталога
type CatalogItem struct {
	ID       int    `json:"id"`
	IblockID int    `json:"iblockId"`
	Name     string `json:"name"`
}
