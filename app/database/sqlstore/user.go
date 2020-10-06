package sqlstore

import (
	dbSql "database/sql"
	"errors"
	"net/http"

	"github.com/netorissi/wk_api_go/entities"
	"gorm.io/gorm"
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
			result.Err = entities.NewAppError("Get", "sqlusersstore", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = user
		}
	})
}

// GetByUniqueFields - Get user for signup
func (sql *SqlUsersStore) GetByUniqueFields(user *entities.User) StoreChannel {
	return Do(func(result *StoreResult) {
		var userExist = &entities.User{}

		response := sql.GetConn().Where(
			"email = ?", user.Email,
		).Or(
			"document = ?", user.Document,
		).Or(
			"cellphone = ?", user.Cellphone,
		).First(&userExist)

		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			var userEmpty *entities.User
			result.Data = userEmpty
		} else if response.Error != nil {
			result.Err = entities.NewAppError("GetByUniqueFields", "sqlusersstore", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = userExist
		}
	})
}

// Create - Create new user
func (sql *SqlUsersStore) Create(user *entities.User) StoreChannel {
	return Do(func(result *StoreResult) {
		response := sql.GetConn().Create(&user)

		if response.Error != nil {
			result.Err = entities.NewAppError("Create", "sqlusersstore", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = user
		}
	})
}

// GetByAuthentication - Get user for login
func (sql *SqlUsersStore) GetByAuthentication(auth *entities.Authentication) StoreChannel {
	return Do(func(result *StoreResult) {
		var user *entities.User

		response := sql.GetConn().Where(
			"email = @access OR document = @access OR cellphone = @access",
			dbSql.Named("access", auth.Access),
		).First(&user)

		if response.Error != nil {
			result.Err = entities.NewAppError("GetByAuthentication", "sqlusersstore", nil, response.Error.Error(), http.StatusInternalServerError)
		} else {
			result.Data = user
		}
	})
}
