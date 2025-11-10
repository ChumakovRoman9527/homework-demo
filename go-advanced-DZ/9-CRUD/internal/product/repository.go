package product

import (
	"9-CRUD_ORDER_API/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepositoryDeps struct {
	DataBase *db.Db
}

type ProductRepository struct {
	DataBase *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		DataBase: database,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.DataBase.Clauses(clause.Returning{}).Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}
