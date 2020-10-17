package app

import (
	"net/http"

	"github.com/netorissi/wk_api_go/entities"
)

// CreateUser - create a new user
func (a *App) CreateUser(user *entities.User) (*entities.User, *entities.AppError) {

	// check exist user
	if userExistErr := a.CheckExistUser(user); userExistErr != nil {
		return nil, userExistErr
	}

	if passwordHash, err := a.HashPassword(user.Password); err != nil {
		return nil, entities.NewAppError("CreateUser", "hashPassword", nil, err.Error(), http.StatusInternalServerError)
	} else {
		user.Password = passwordHash
	}

	result := <-a.Srv.SqlStore.Users().Create(user)
	if result.Err != nil {
		return nil, result.Err
	}

	return result.Data.(*entities.User), nil
}

// GetUserByAuthentication - get user to authentication
func (a *App) GetUserByAuthentication(auth *entities.Authentication) (
	*entities.ResponseAuth,
	*entities.AppError,
) {
	result := <-a.Srv.SqlStore.Users().GetByAuthentication(auth)

	if result.Err != nil {
		return nil, result.Err
	}

	user := result.Data.(*entities.User)

	responseAuth := &entities.ResponseAuth{
		User:  user,
		Token: "",
	}

	return responseAuth, nil
}

// CheckExistUser - Verify exist user in database
func (a *App) CheckExistUser(user *entities.User) *entities.AppError {

	result := <-a.Srv.SqlStore.Users().GetByUniqueFields(user)
	if result.Err != nil {
		return result.Err
	}

	userExist := result.Data.(*entities.User)
	if userExist != nil {
		return entities.NewAppError("CheckExistUser", "getbyuniquefields", nil, "User existing to platform!", http.StatusBadRequest)
	}

	return nil
}
