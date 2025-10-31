package link

import "9-CRUD/pkg/db"

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

func (repo *LinkRepository) Create(link *Link) {

}
