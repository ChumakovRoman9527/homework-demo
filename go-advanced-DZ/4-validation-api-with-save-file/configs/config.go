package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailValidationConfig struct {
	EmailConfig Email
}
type Email struct {
	Email       string
	Password    string
	Address     string
	Port        string
	Hash_secret string
}

func LoadConfig() *EmailValidationConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла, используем конфигурацию по умолчанию")
	}

	return &EmailValidationConfig{
		EmailConfig: Email{
			Email:       os.Getenv("Email"),
			Password:    os.Getenv("Password"),
			Address:     os.Getenv("Address"),
			Port:        os.Getenv("Port"),
			Hash_secret: os.Getenv("HASH_SECRET"),
		},
	}
}
