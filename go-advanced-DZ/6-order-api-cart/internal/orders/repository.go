package orders

import (
	"6-order-api-cart/internal/product"
	"6-order-api-cart/pkg/db"
	"fmt"
)

type OrderRepositoryDeps struct {
	DataBase *db.Db
}

type OrderRepository struct {
	DataBase *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		DataBase: database,
	}
}

func (repo *OrderRepository) CreateOrder(UserID int, phone string, newOrder CreateOrderRequest) (Order, error) {
	tx := repo.DataBase.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	NewDBOrder := Order{UserID: UserID, UserPhone: phone, OrderStatus: OrderStatusCreated}
	result := tx.Create(&NewDBOrder)
	if err := result.Error; err != nil {
		tx.Error = err
		return Order{}, result.Error
	}
	OrderId := NewDBOrder.ID
	var NewOrderDetails []OrderDetails
	for _, value := range newOrder.Items {
		NewOrderDetails = append(NewOrderDetails,
			OrderDetails{OrderID: OrderId, ProductID: uint(value.ProductID), Quantity: value.Quantity})
	}

	result = tx.Create(&NewOrderDetails)
	if err := result.Error; err != nil {
		tx.Error = err
		return Order{}, result.Error
	}

	NewDBOrder.Items = NewOrderDetails

	return NewDBOrder, nil
}

func (repo *OrderRepository) GetOrder(OrderId int) (GetOrderResponse, error) {
	var OrderDB Order
	var Order GetOrderResponse
	var OrderItems []ProductResponse

	// var product product.Product

	repo.DataBase.Find(&OrderDB, OrderId)
	Order.OrderID = int(OrderDB.ID)
	Order.Created_At = OrderDB.CreatedAt.Format("yyyy-MM-dd")
	Order.OrderStatus = OrderDB.OrderStatus

	for _, value := range OrderDB.Items {
		OrderItems = append(OrderItems, ProductResponse{
			Product: product.Product{},
			//Product:  product.ProductRepository.GetByID(value.ProductID),
			Quantity: value.Quantity,
		})
		fmt.Println(value.ProductID)
	}

	return Order, nil
}
