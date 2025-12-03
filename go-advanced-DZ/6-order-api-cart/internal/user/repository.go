package user

import (
	"6-order-api-cart/pkg/db"
	"fmt"

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

// func (repo *UserRepository) Create(Phone string) (*User, error) {
// 	newUser := User{
// 		Phone: Phone,
// 	}
// 	result := repo.DataBase.Clauses(clause.Returning{}).Create(newUser)
// 	fmt.Println("User created", result)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &newUser, nil
// }

func (repo *UserRepository) GetIdByPhone(Phone string) (int, error) {
	var _user User
	db_result := repo.DataBase.Where(&User{Phone: Phone}).First(&_user)
	fmt.Println(db_result)
	fmt.Println(_user)
	// if db_result.Error != nil {
	// 	return -1, db_result.Error
	// }
	if _user.ID == 0 {
		_user = User{Phone: Phone}
		fmt.Println("2 user", _user)
		db_result = repo.DataBase.Clauses(clause.Returning{}).Create(&_user)
		if db_result.Error != nil {
			return -1, db_result.Error
		}

	}
	res := int(_user.ID)
	return res, nil
}
