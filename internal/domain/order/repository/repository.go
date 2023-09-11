package repository

import (
	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/infras"
)

type OrderRepository interface {
	BlibliOrdersRepository
	ShopeeOrdersRepository
}

type OrderRepositoryMySQL struct {
	DB  *infras.MySQLConn
	CFG *configs.Config
}

func ProvideOrderRepositoryMySQL(db *infras.MySQLConn, cfg *configs.Config) *OrderRepositoryMySQL {
	return &OrderRepositoryMySQL{
		DB:  db,
		CFG: cfg,
	}
}
