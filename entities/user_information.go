package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInformation struct {
	ID               string `gorm:"type:char(36);default:(HEX(UUID_TO_BIN(UUID())));" json:"id"`
	UserID           string `gorm:"type:char(36);default:'';not null"`
	EmailConfirm     string `gorm:"int(1);default:0" json:"email_confirm"`
	EmailUpdated     int64  `gorm:"autoUpdateTime;default:0" json:"email_updated"`
	PasswordUpdated  int64  `gorm:"autoUpdateTime;default:0" json:"password_updated"`
	CellphoneUpdated int64  `gorm:"autoUpdateTime;default:0" json:"cellphone_updated"`
	Updated          int64  `gorm:"autoUpdateTime;default:0" json:"updated"`
	Created          int64  `gorm:"autoCreateTime;default:0" json:"created"`
}

// BeforeCreate will set a UUID v4 rather than numeric ID.
func (u *UserInformation) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
