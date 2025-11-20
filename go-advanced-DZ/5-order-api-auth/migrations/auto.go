package main

import (
	"5-order-api-auth/internal/auth"
	"5-order-api-auth/internal/link"
	"5-order-api-auth/internal/product"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&link.Link{}, &product.Product{}, &auth.PhoneAuth{})
	if err != nil {
		panic(err)
	}
}
