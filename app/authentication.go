package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/netorissi/wk_api_go/entities"
	"github.com/netorissi/wk_api_go/utils"
)

func (a *App) Login(auth *entities.Authentication, r *http.Request) (
	*entities.ResponseAuth,
	*entities.AppError,
) {
	// validar dados

	var responseAuth *entities.ResponseAuth
	var err *entities.AppError
	var token string

	if responseAuth, err = a.GetUserByAuthentication(auth); err != nil || responseAuth == nil {
		return nil, entities.NewAppError("Login", "[COD-AUTH-4]", nil, err.ToJson(), http.StatusUnauthorized)
	}

	if token, err = a.CreateToken(responseAuth.User); err != nil {
		return nil, entities.NewAppError("Login", "[COD-AUTH-5]", nil, err.ToJson(), http.StatusInternalServerError)
	}

	preSession := &entities.Session{
		DeviceID: auth.DeviceID,
		Token:    token,
		UserID:   responseAuth.User.ID,
	}

	if _, err = a.CreateSession(r, preSession); err != nil {
		return nil, entities.NewAppError("Login", "[COD-AUTH-6]", nil, err.ToJson(), http.StatusInternalServerError)
	}

	responseAuth.Token = token

	return responseAuth, nil
}

func (a *App) Logout() {}

func (a *App) CreateToken(user *entities.User) (string, *entities.AppError) {
	secret := os.Getenv(utils.KEY_SECRET)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
		"iat":   time.Now().Unix(),
	})

	tokenSecret, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", entities.NewAppError("CreateToken", "[COD-AUTH-0]", nil, err.Error(), http.StatusBadRequest)
	}

	return tokenSecret, nil
}

func (a *App) RefreshToken() {}

func (a *App) ValidToken(token string) (bool, *entities.AppError) {
	if token == "" {
		return false, entities.NewAppError("ValidToken", "[COD-AUTH-1]", nil, "Invalid token.", http.StatusUnauthorized)
	}

	authToken, err := a.StringToToken(token)
	if err != nil {
		return false, err
	}

	if !authToken.Valid {
		return false, entities.NewAppError("ValidToken", "[COD-AUTH-2]", nil, "Invalid token.", http.StatusUnauthorized)
	}

	return true, nil
}

func (a *App) StringToToken(tokenString string) (*jwt.Token, *entities.AppError) {
	secret := os.Getenv(utils.KEY_SECRET)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, entities.NewAppError("StringToToken", "[COD-AUTH-3]", nil, "Permission denied.", http.StatusUnauthorized)
	}

	return token, nil
}
