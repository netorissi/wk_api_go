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
