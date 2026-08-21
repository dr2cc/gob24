package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	// ENV=local используется для разработки на компьютере.
	// На реальном сервере значение "local" перезапишется реальным значением переменной ENV (например, "prod").
	// Точно также его использует логгер- от максимално разверноутого лога при разработке, до строго json на сервере
	Env        string `env:"ENV" envDefault:"local"`
	WebhookURL string `env:"B24_WEBHOOK_URL,required"`
	UserPath   string `env:"USER_PATH"`
	// BaseURL    string `env:"BASE_URL"`
	// CacheDumpPath string `env:"FILE_STORAGE_PATH"`
}

func New() (*Config, error) {
	// Создаем пустую структуру (выделяем память)
	cfg := &Config{}

	// // 1. Присваиваем полям Config флаги
	// // Разбираем флаги в конфигурацию.
	// flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "base URL")
	// flag.StringVar(&cfg.CacheDumpPath, "f", "./short-url-db.json", "path to dump with addresses")
	// flag.Parse()

	// 2. Загужаем в окружение переменные из файла .env в корне.
	if err := godotenv.Load(); err != nil {
		// // В теории нужно писать что-то типа:
		// return nil, fmt.Errorf("failed to load .env file: %w", err)
		// // Но правильнее оставлять только лог, так как в Docker .env файла может не быть
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	// 3. parse читает переменные окружения и заполняет структуру cfg .
	// Подробнее:
	// - библиотека берёт структуру cfg и с помощью механизма рефлексии (способность программы изучать саму себя) смотрит на её поля;
	// - она ищет поля, у которых есть тег env:"ИМЯ_ПЕРЕМЕННОЙ";
	// - если в окружении ОС есть переменная с таким же именем библиотека caarlos0/env берёт её значение и записывает в это конкретное поле.
	// Библиотека сама проверит тег "required" и вернет ошибку, если переменной c таким тегом нет.
	if err := env.Parse(cfg); err != nil {
		// log.Fatalf("Config parsing error: %+v\n", err)
		return nil, fmt.Errorf("unable to parse env variables: %w", err)
	}

	return cfg, nil
}
