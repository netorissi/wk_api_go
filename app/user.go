package app

import "github.com/netorissi/wk_api_go/entities"

// CreateUser - create a new user
func (a *App) CreateUser(user *entities.User) (*entities.User, *entities.AppError) {
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
