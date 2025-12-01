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
