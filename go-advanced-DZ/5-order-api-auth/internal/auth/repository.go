package auth

import (
	"5-order-api-auth/pkg/db"

	"gorm.io/gorm/clause"
)

type PhoneAuthRepositoryDeps struct {
	DataBase *db.Db
}

type PhoneAuthRepository struct {
	DataBase *db.Db
}

func NewPhoneAuthRepository(database *db.Db) *PhoneAuthRepository {
	return &PhoneAuthRepository{
		DataBase: database,
	}
}

func (repo *PhoneAuthRepository) Create(phoneAuth *PhoneAuth) (*PhoneAuth, error) {
	result := repo.DataBase.Clauses(clause.Returning{}).Create(phoneAuth)
	if result.Error != nil {
		return nil, result.Error
	}
	return phoneAuth, nil
}

func (repo *PhoneAuthRepository) GetBySessionCode(sessionId string, code string) (*PhoneAuth, error) {
	var phoneAuth PhoneAuth
	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	result := repo.DataBase.Where(&PhoneAuth{SessionID: sessionId, Code: code}).First(&phoneAuth)
	if result.Error != nil {
		return nil, result.Error
	}
	return &phoneAuth, nil
}

func (repo *PhoneAuthRepository) DeleteBySessionCode(sessionId string, code string) (*PhoneAuth, error) {
	var phoneAuth PhoneAuth
	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	result := repo.DataBase.Delete(&PhoneAuth{SessionID: sessionId, Code: code})
	if result.Error != nil {
		return nil, result.Error
	}
	return &phoneAuth, nil
}
