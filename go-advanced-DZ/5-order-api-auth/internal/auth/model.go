package auth

import (
	"gorm.io/gorm"
)

type PhoneAuth struct {
	gorm.Model
	Phone     string
	SessionID string
	Code      string
}
