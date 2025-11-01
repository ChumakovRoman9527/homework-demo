package link

import (
	"9-CRUD/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepositoryDeps struct {
	DataBase *db.Db
}

type LinkRepository struct {
	DataBase *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		DataBase: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.DataBase.DB.First(&link, "hash=?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) CheckByHash(hash string) (bool, error) {
	var link Link
	var exists bool
	result := repo.DataBase.DB.Where(&link, "hash=?", hash).Find(&exists)
	if result.Error != nil {
		return false, result.Error
	}
	return exists, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Delete(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
