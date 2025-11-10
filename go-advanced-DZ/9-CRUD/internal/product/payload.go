package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string         `json:"Name" validate:"required"`
	Description string         `json:"Description" validate:"required"`
	Images      pq.StringArray `json:"Image" gorm:"type:text[]"`
}

type ProductUpdateRequest struct {
	ID          uint           `json:"Id" validate:"required, number"`
	Name        string         `json:"Name" validate:"required"`
	Description string         `json:"Description" validate:"required"`
	Images      pq.StringArray `json:"Image" gorm:"type:text[]"`
}
