package auth

import (
	"gorm.io/gorm"
)

type PhoneAuth struct {
	gorm.Model
	Phone     string
	UserId    uint
	SessionID string
	Code      string
}
