package user

import (
	"11-JWTAUTH/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepositoryDeps struct {
	DataBase *db.Db
}

type UserRepository struct {
	DataBase *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		DataBase: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.DataBase.Where("email =", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
