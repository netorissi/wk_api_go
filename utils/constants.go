package utils

import "os"

const KEY_SECRET = "WK_SECRET"

func SetSecretAuth() {
	os.Setenv(KEY_SECRET, "5a867a27ad5f178545aa15146d0df613")
}

const (
	RoleUser     = "system_user"
	RoleProvider = "provider"
)
