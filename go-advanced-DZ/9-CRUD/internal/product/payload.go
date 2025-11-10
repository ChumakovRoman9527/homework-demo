package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string         `json:"Name" validate:"required"`
	Description string         `json:"Description" validate:"required"`
	Images      pq.StringArray `json:"Image" gorm:"type:text[]"`
}

type ProductUpdateRequest struct {
	Name        string         `json:"Name"`
	Description string         `json:"Description"`
	Images      pq.StringArray `json:"Image" gorm:"type:text[]"`
}
