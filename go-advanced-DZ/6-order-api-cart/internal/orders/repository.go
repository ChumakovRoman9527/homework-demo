package orders

import (
	"6-order-api-cart/internal/product"
	"6-order-api-cart/pkg/db"
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

func (repo *OrderRepository) GetOrder(OrderId uint64) (GetOrderResponse, error) {
	var OrderDB Order
	var Order GetOrderResponse
	var OrderItems []OrderDetails

	// var product product.Product

	repo.DataBase.Find(&OrderDB, OrderId)
	Order.OrderID = int(OrderDB.ID)
	Order.Created_At = OrderDB.CreatedAt.Format("2006-01-02 15:04:05")
	Order.OrderStatus = OrderDB.OrderStatus

	repo.DataBase.Find(&OrderItems, "order_id = ?", OrderId)

	for _, value := range OrderItems { //можно попробовать паралелить запросы к БД
		var product product.Product
		var orderProduct ProductResponse
		repo.DataBase.Find(&product, "id=?", value.ProductID)
		//product = product.Product.GetByID(value.ID)
		orderProduct = ProductResponse{Quantity: value.Quantity,
			Product: product}

		Order.Products = append(Order.Products, orderProduct)
	}

	return Order, nil
}

func (repo *OrderRepository) GetUserOrders(UserID uint64) (Orders, error) {
	var OrderDB []Order
	var UserOrders Orders
	repo.DataBase.Find(&OrderDB, "user_id = ? ", UserID)
	for _, order := range OrderDB { //можно попробовать паралелить запросы к БД
		var Order GetOrderResponse
		Order, _ = repo.GetOrder(uint64(order.ID))

		UserOrders.Orders = append(UserOrders.Orders, Order)
	}

	return UserOrders, nil
}
