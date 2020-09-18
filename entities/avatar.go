package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Avatar struct {
	ID      string `gorm:"type:char(36);default:(HEX(UUID_TO_BIN(UUID())));" json:"id"`
	Name    string `gorm:"type:varchar(200);default:''" json:"name"`
	Type    string `gorm:"type:int(1);default:0" json:"type"`
	URL     string `gorm:"type:varchar(255);default:'';not null" json:"url"`
	Updated int64  `gorm:"autoUpdateTime;default:0" json:"updated"`
	Created int64  `gorm:"autoCreateTime;default:0" json:"created"`
}

// BeforeCreate will set a UUID v4 rather than numeric ID.
func (a *Avatar) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return
}
