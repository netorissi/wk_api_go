package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/netorissi/wk_api_go/entities"
	"github.com/netorissi/wk_api_go/utils"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) Login(auth *entities.Authentication, r *http.Request) (
	*entities.ResponseAuth,
	*entities.AppError,
) {
	// validar dados

	var responseAuth *entities.ResponseAuth
	var err *entities.AppError
	var token string
	var refreshToken string

	if responseAuth, err = a.GetUserByAuthentication(auth); err != nil || responseAuth == nil {
		return nil, err
	}

	if token, err = a.CreateToken(responseAuth.User); err != nil {
		return nil, err
	}

	if refreshToken, err = a.RefreshToken(); err != nil {
		return nil, err
	}

	preSession := &entities.Session{
		DeviceID: auth.DeviceID,
		Token:    token,
		UserID:   responseAuth.User.ID,
	}

	if _, err = a.CreateSession(r, preSession); err != nil {
		return nil, err
	}

	responseAuth.Token = token
	responseAuth.RefreshToken = refreshToken

	return responseAuth, nil
}

func (a *App) Logout() {}

func (a *App) CreateToken(user *entities.User) (string, *entities.AppError) {
	secret := os.Getenv(utils.KEY_SECRET)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
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

func (a *App) RefreshToken() (string, *entities.AppError) {
	secret := os.Getenv(utils.KEY_SECRET)

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshTokenSecret, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", entities.NewAppError("RefreshToken", "[COD-AUTH-4]", nil, err.Error(), http.StatusBadRequest)
	}

	return refreshTokenSecret, nil
}

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

// HashPassword - generate hash by password
func (a *App) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - compare password to hash
func (a *App) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
