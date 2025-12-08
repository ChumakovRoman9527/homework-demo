package main

import (
	"6-order-api-cart/internal/auth"
	"6-order-api-cart/internal/link"
	"6-order-api-cart/internal/orders"
	"6-order-api-cart/internal/product"
	"6-order-api-cart/internal/user"
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
	err = db.AutoMigrate(&link.Link{}, &product.Product{}, &auth.PhoneAuth{}, orders.Order{}, orders.OrderDetails{}, user.User{})
	if err != nil {
		panic(err)
	}
}
