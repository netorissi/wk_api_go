package entities

import (
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
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
