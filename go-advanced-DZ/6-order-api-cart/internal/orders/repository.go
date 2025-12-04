package orders

import "6-order-api-cart/pkg/db"

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

func (repo *OrderRepository) CreateOrder(UserID int, phone string, newOrder *CreateOrderRequest) (Order, error) {
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
	var NewDBOrder Order
	NewDBOrder = Order{UserID: UserID, UserPhone: phone, OrderStatus: OrderStatusCreated}
	result := tx.Create(&NewDBOrder)
	if err := result.Error; err != nil {
		tx.Error = err
		return Order{}, result.Error
	}
	OrderId := NewDBOrder.ID

	OrderDetails := OrderDetails{OrderID: OrderId, ProductID: uint(newOrder.ProductID), Quantity: newOrder.Quantity}
	result = tx.Create(&OrderDetails)
	if err := result.Error; err != nil {
		tx.Error = err
		return Order{}, result.Error
	}

	return NewDBOrder, nil
}
