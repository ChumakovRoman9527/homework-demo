package main

import (
	"6-order-api-cart/internal/auth"
	"6-order-api-cart/internal/orders"
	"6-order-api-cart/internal/product"
	"6-order-api-cart/internal/user"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	products := []product.Product{
		{Name: "Test1"},
		{Name: "Test2"},
		{Name: "Test3"},
		{Name: "Test4"},
		{Name: "Test5"},
		{Name: "Test6"},
		{Name: "Test7"},
		{Name: "Test8"},
		{Name: "Test9"},
		{Name: "Test10"},
		{Name: "Test11"}}
	db.Create(&products)
}
func removeData(db *gorm.DB) {
	db.Unscoped().Where("phone = ?", "89123456789").Delete(&user.User{})
	db.Unscoped().Where("phone = ?", "89123456789").Delete(&auth.PhoneAuth{})
	db.Unscoped().Where("1 = ?", "1").Delete(&orders.OrderDetails{})
	db.Unscoped().Where("1 = ?", "1").Delete(&orders.Order{})
	// db.Unscoped().Where("1 = ?", "1").Delete(&product.Product{})
}

func TestOrderSuccess(t *testing.T) {
	//Prepare
	db := initDB()
	initData(db)
	//Test
	ts := httptest.NewServer(App())
	defer removeData(db)
	defer ts.Close()
	//1. стучимся по номеру телефона
	//2. отправляем полученный код из смс и получаем токен
	//3. формируем заказ и отправляем

	data, _ := json.Marshal(&auth.LoginPhoneRequest{
		Phone: "89123456789",
	})

	res, err := http.Post(ts.URL+"/auth/loginphone", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var phoneResponse auth.LoginPhoneResponse

	err = json.Unmarshal(body, &phoneResponse)
	if err != nil {
		t.Fatal(err)
	}

	if phoneResponse.SessionID == "" {
		t.Fatal("session nil !")
	}

	//тут находим код СМС из тестовой БД
	// SMSCode := db.Select("code").Where("phone = ?", "89123456789").First(&auth.PhoneAuth{})
	var SMSCode string
	db.Table("phone_auths").Select("code").Where("phone = ?", "89123456789").Take(&SMSCode)
	data, _ = json.Marshal(&auth.LoginSMSRequest{
		LoginPhoneResponse: phoneResponse,
		Code:               SMSCode,
	})

	res, err = http.Post(ts.URL+"/auth/loginSMS", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var SMSResponse auth.LoginSMSResponse

	err = json.Unmarshal(body, &SMSResponse)
	if err != nil {
		t.Fatal(err)
	}

	if SMSResponse.TOKEN == "" {
		t.Fatal("TOKEN nil !")
	}

	Bearer := "Bearer " + SMSResponse.TOKEN

	order_items := []orders.ItemOrder{
		{ProductID: 1,
			Quantity: 10},
		{ProductID: 2,
			Quantity: 10},
		{ProductID: 3,
			Quantity: 10},
		{ProductID: 4,
			Quantity: 10},
		{ProductID: 5,
			Quantity: 10},
		{ProductID: 6,
			Quantity: 10},
		{ProductID: 7,
			Quantity: 10},
	}

	data, _ = json.Marshal(&orders.CreateOrderRequest{order_items})

	req, err := http.NewRequest("POST", ts.URL+"/order", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}
	req.Header.Set("Authorization", Bearer)

	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Expected %d got %d", http.StatusCreated, res.StatusCode)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var OrderResponse orders.GetOrderResponse

	err = json.Unmarshal(body, &OrderResponse)
	if err != nil {
		t.Fatal(err)
	}
}
