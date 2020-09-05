package sqlstore

import (
	"github.com/netorissi/wk_api_go/app/infra"
	"github.com/netorissi/wk_api_go/entities"
	"gorm.io/gorm"
)

type SqlSupplier struct {
	connection *gorm.DB
	oldStores  SqlSupplierOldStores
	settings   *entities.Config
}

type SqlSupplierOldStores struct{}

func NewSqlSupplier(configs *entities.Config) *SqlSupplier {
	supplier := &SqlSupplier{
		settings: configs,
	}

	supplier.initConnection()

	return supplier
}

func (s *SqlSupplier) initConnection() {
	s.connection = infra.OpenConnectionMySQL(s.settings.Urls.MySQL)
}

func (s *SqlSupplier) GetConn() *gorm.DB {
	return s.connection
}

func (s *SqlSupplier) Close() {
	if s.connection != nil {
		sqlDB, err := s.connection.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}
