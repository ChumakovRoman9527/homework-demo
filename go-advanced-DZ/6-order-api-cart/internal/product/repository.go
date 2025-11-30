package product

import (
	"6-order-api-cart/pkg/db"

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

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetByID(id uint) (*Product, error) {
	var product Product
	result := repo.DataBase.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.DataBase.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
