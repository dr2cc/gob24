package app

import (
	"github.com/dr2cc/gob24/internal/bitrix/user"
)

func Run() error {
	// // Сейчас (17.08.2026) создаю client и ctx непосредственно в catalog.Catalog()
	// client := b24.NewClient(os.Getenv("B24_WEBHOOK_URL"))
	// ctx := context.Background()

	// catalog.Catalog()
	user.UserAdd()

	return nil
}
