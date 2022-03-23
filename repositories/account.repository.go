package repositories

import "github.com/betaapskaita/beta-server/libs/databases"

type AccountRepository struct {
	db databases.Database
}

func NewAccountRepository(db databases.Database) AccountRepository {
	return AccountRepository{db}
}
