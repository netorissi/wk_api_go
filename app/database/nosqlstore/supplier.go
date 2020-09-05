package nosqlstore

import (
	"context"
	"time"

	"github.com/netorissi/wk_api_go/app/infra"
	"github.com/netorissi/wk_api_go/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MAX_DB_CONN_LIFETIME    = 60
	DB_PING_ATTEMPTS        = 18
	DB_PING_TIMEOUT_SECS    = 10
	DB_CONTEXT_TIMEOUT_SECS = 10
)

type NoSqlStore interface {
	GetConn() *mongo.Database
	Close()
	GetCtx() (context.Context, context.CancelFunc)
}

type NoSqlSupplierOldStores struct{}

type NoSqlSupplier struct {
	connection *mongo.Client
	oldStores  NoSqlSupplierOldStores
	settings   *entities.Config
	ctx        context.Context
}

func NewNoSqlSupplier(configs *entities.Config) *NoSqlSupplier {
	supplier := &NoSqlSupplier{
		settings: configs,
	}

	supplier.initConnection()

	return supplier
}

func (s *NoSqlSupplier) initConnection() {
	s.connection, s.ctx = infra.OpenConnectionNoSQL(s.settings.Urls.NoSQL)
}

func (s *NoSqlSupplier) GetConn() *mongo.Database {
	return s.connection.Database(s.settings.Urls.NoSQL)
}

func (s *NoSqlSupplier) Close() {
	if s.connection != nil {
		s.connection.Disconnect(s.ctx)
	}
}

func (s *NoSqlSupplier) GetCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DB_CONTEXT_TIMEOUT_SECS*time.Second)
}
