package orders

import (
	"6-order-api-cart/internal/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	//ниже закоментировал, так как нельзя связать логин пользователя через jwt с заказом, это все таки логин.  логин должен давать ИД пользователя для которого и вытаскивать всю инфу
	//похорошему надо авторизацию связывать с пользователями. так будет правильнее
	// PhoneAuthID uint           `gorm:"index;not null"` // Foreign key
	// PhoneAuth   auth.PhoneAuth `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserPhone   string `json:"phone" gorm:"index"`
	OrderStatus string `json:"status"`
	//	OrderAddress string         `json:"address"` это свойство пользователя все таки и это свойство доставки а не заказа, одна доставка - множество заказов, и их может быть много !
	Items []OrderDetails `gorm:"foreignKey:OrderID;references:ID"`
}

type OrderDetails struct {
	gorm.Model
	OrderID   uint            `gorm:"index;not null"`
	ProductID uint            `gorm:"index;not null"`
	Product   product.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Quantity  int
}
