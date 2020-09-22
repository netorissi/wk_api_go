package sqlstore

import (
	"net/http"

	"github.com/netorissi/wk_api_go/entities"
)

type SqlSessionsStore struct {
	SqlStore
}

func NewSqlSessionsStore(sqlStore SqlStore) SessionsStore {
	sql := &SqlSessionsStore{
		SqlStore: sqlStore,
	}

	return sql
}

// Get - Get session
func (sql *SqlSessionsStore) Get(session *entities.Session) StoreChannel {
	return Do(func(result *StoreResult) {
		response := sql.GetConn().First(&session, session.Token)

		if response.Error != nil {
			result.Err = entities.NewAppError("CreateSession", "sqlsessionsstore.create", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = session
		}
	})
}

// Create - Create new session
func (sql *SqlSessionsStore) Create(session *entities.Session) StoreChannel {
	return Do(func(result *StoreResult) {
		response := sql.GetConn().Create(&session)

		if response.Error != nil {
			result.Err = entities.NewAppError("CreateSession", "sqlsessionsstore.create", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = session
		}
	})
}
