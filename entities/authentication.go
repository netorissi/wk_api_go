package entities

import (
	"encoding/json"
	"io"
)

// Authentication - data to login
type Authentication struct {
	Access   string
	Password string
	DeviceID string
}

// ResponseAuth - response to user after login
type ResponseAuth struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// AuthenticationFromJSON - convert to struct
func AuthenticationFromJSON(data io.Reader) *Authentication {
	decoder := json.NewDecoder(data)
	var auth *Authentication

	if err := decoder.Decode(&auth); err == nil {
		return auth
	}

	return nil
}

// ResponseAuthToJSON - convert to json
func (u *ResponseAuth) ResponseAuthToJSON() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(b)
}
