package entities

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:char(36);default:(HEX(UUID_TO_BIN(UUID())));" json:"id"`
	Name      string `gorm:"type:varchar(200);default:''" json:"name"`
	Document  string `gorm:"type:varchar(500);default:''" json:"document"`
	Password  string `gorm:"type:varchar(100);default:''" json:"password"`
	Email     string `gorm:"type:varchar(500);default:''" json:"email"`
	Status    int    `gorm:"type:int(1);default:1" json:"status"`
	Roles     string `gorm:"type:varchar(255);default:''" json:"roles"`
	Cellphone string `gorm:"type:varchar(20);default:''" json:"cellphone"`
	Bio       string `gorm:"type:varchar(500);default:''" json:"bio"`
	Provider  int    `gorm:"type:int(1);default:0" json:"provider"`
	AvatarID  string `gorm:"type:char(36);default:'';not null"`
	Avatar    Avatar `gorm:"foreignKey:AvatarID;" json:"avatar"`
	Updated   int64  `gorm:"autoUpdateTime;default:0" json:"updated"`
	Created   int64  `gorm:"autoCreateTime;default:0" json:"created"`
}

// BeforeCreate will set a UUID v4 rather than numeric ID.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.Email = strings.ToLower(u.Email)

	return nil
}

// BeforeUpdate check fields requireds.
func (u *User) BeforeUpdate(tx *gorm.DB) error {

	if len(u.ID) == 0 {
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
