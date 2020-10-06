package entities

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"type:int;auto_increment;primary_key;" json:"id"`
	Name      string `gorm:"type:varchar(200);default:''" json:"name"`
	Document  string `gorm:"type:varchar(500);default:''" json:"document"`
	Password  string `gorm:"type:varchar(100);default:''" json:"password"`
	Email     string `gorm:"type:varchar(500);default:''" json:"email"`
	Status    int    `gorm:"type:int(1);default:1" json:"status"`
	Roles     string `gorm:"type:varchar(255);default:''" json:"roles"`
	Cellphone string `gorm:"type:varchar(20);default:''" json:"cellphone"`
	Bio       string `gorm:"type:varchar(500);default:''" json:"bio"`
	AvatarID  int    `gorm:"type:int;"`
	Avatar    Avatar `gorm:"foreignKey:AvatarID;" json:"avatar"`
	Updated   int64  `gorm:"autoUpdateTime;default:0" json:"updated"`
	Created   int64  `gorm:"autoCreateTime;default:0" json:"created"`
}

// BeforeCreate will set a UUID v4 rather than numeric ID.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Email = strings.ToLower(u.Email)
	u.AvatarID = 1
	return nil
}

// BeforeUpdate check fields requireds.
func (u *User) BeforeUpdate(tx *gorm.DB) error {

	if u.ID == 0 {
		return errors.New("User don't have ID")
	}

	u.Email = strings.ToLower(u.Email)

	return nil
}

// UserFromJSON - convert to struct
func UserFromJSON(data io.Reader) *User {
	decoder := json.NewDecoder(data)
	var user *User

	if err := decoder.Decode(&user); err == nil {
		return user
	}

	return nil
}

// UserToJSON - convert to json
func (u *User) UserToJSON() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(b)
}
