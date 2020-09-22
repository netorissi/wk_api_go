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

type SqlStore interface {
	GetConn() *gorm.DB
	Close()

	Users() UsersStore
	Sessions() SessionsStore
}

type SqlSupplierOldStores struct {
	users    UsersStore
	sessions SessionsStore
}

func NewSqlSupplier(configs *entities.Config) *SqlSupplier {
	supplier := &SqlSupplier{
		settings: configs,
	}

	supplier.initConnection()

	supplier.oldStores.users = NewSqlUsersStore(supplier)
	supplier.oldStores.sessions = NewSqlSessionsStore(supplier)

	return supplier
}

func (s *SqlSupplier) initConnection() {
	s.connection = infra.OpenConnectionMySQL(s.settings.Urls.MySQL)
	s.Migrate()
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

func (ss *SqlSupplier) Users() UsersStore {
	return ss.oldStores.users
}

func (ss *SqlSupplier) Sessions() SessionsStore {
	return ss.oldStores.sessions
}
