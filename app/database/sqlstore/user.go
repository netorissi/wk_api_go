package sqlstore

import (
	"net/http"

	"github.com/netorissi/wk_api_go/entities"
)

type SqlUsersStore struct {
	SqlStore
}

func NewSqlUsersStore(sqlStore SqlStore) UsersStore {
	sql := &SqlUsersStore{
		SqlStore: sqlStore,
	}

	return sql
}

// Get - Get user
func (sql *SqlUsersStore) Get(user *entities.User) StoreChannel {
	return Do(func(result *StoreResult) {
		response := sql.GetConn().First(&user, user.ID)

		if response.Error != nil {
			result.Err = entities.NewAppError("CreateSession", "sqlsessionsstore.create", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = user
		}
	})
}

// Create - Create new user
func (sql *SqlUsersStore) Create(user *entities.User) StoreChannel {
	return Do(func(result *StoreResult) {
		response := sql.GetConn().Create(&user)

		if response.Error != nil {
			result.Err = entities.NewAppError("CreateUser", "sqlusersstore.create", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = user
		}
	})
}

// GetByAuthentication - Get user for login
func (sql *SqlUsersStore) GetByAuthentication(auth *entities.Authentication) StoreChannel {
	return Do(func(result *StoreResult) {
		var user *entities.User

		response := sql.GetConn().First(&user, user.ID)

		if response.Error != nil {
			result.Err = entities.NewAppError("CreateSession", "sqlsessionsstore.create", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = user
		}
	})
}
