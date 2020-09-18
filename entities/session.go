package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID        string `gorm:"type:char(36);default:(HEX(UUID_TO_BIN(UUID())));" json:"id"`
	UserID    string `gorm:"type:char(36);default:'';not null"`
	DeviceID  string `gorm:"type:varchar(200);default:''" json:"device_id"`
	Token     string `gorm:"type:varchar(100);default:''" json:"token"`
	Status    int    `gorm:"type:int(1);default:1" json:"status"`
	IPAddress string `gorm:"type:varchar(30);default:''" json:"ip_address"`
	Expires   int64  `gorm:"type:bigint;default:0" json:"expires"`
	Updated   int64  `gorm:"autoUpdateTime;default:0" json:"updated"`
	Created   int64  `gorm:"autoCreateTime;default:0" json:"created"`
}

// BeforeCreate will set a UUID v4 rather than numeric ID.
func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}
