package orders

import "6-order-api-cart/internal/product"

type CreateOrderRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type GetOrderResponse struct {
	OrderID     int               `json:"order_id"`
	Created_At  string            `json:"order_date"`
	OrderStatus string            `json:"order_status"`
	Products    []ProductResponse `json:"products"`
}

type ProductResponse struct {
	product.Product
}

type Orders struct {
	Orders []GetOrderResponse `json:"orders"`
}
