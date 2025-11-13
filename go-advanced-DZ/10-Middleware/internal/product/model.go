package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"Name"`
	Description string         `json:"Description"`
	Images      pq.StringArray `json:"Image" gorm:"type:text[]"`
}

func newProduct(newProduct ProductCreateRequest) *Product {
	product := &Product{
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Images:      newProduct.Images,
	}

	return product
}
