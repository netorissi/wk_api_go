package entities

import (
	"github.com/netorissi/wk_api_go/utils"
	"gorm.io/gorm"
)

type Session struct {
	ID        int    `gorm:"type:int;auto_increment; primary_key;" json:"id"`
	UserID    int    `gorm:"type:int;not null"`
	User      User   `gorm:"foreignKey:UserID;" json:"user"`
	DeviceID  string `gorm:"type:varchar(200);default:''" json:"device_id"`
	Token     string `gorm:"type:varchar(255);default:''" json:"token"`
	Status    int    `gorm:"type:int(1);default:1" json:"status"`
	IPAddress string `gorm:"type:varchar(30);default:''" json:"ip_address"`
	Expires   int64  `gorm:"type:bigint;default:0" json:"expires"`
	Updated   int64  `gorm:"autoUpdateTime;default:0" json:"updated"`
	Created   int64  `gorm:"autoCreateTime;default:0" json:"created"`
}

// BeforeCreate will set created.
func (u *Session) BeforeCreate(tx *gorm.DB) error {
	u.Created = utils.DateNow()
	return nil
}

// BeforeUpdate check fields requireds.
func (u *Session) BeforeUpdate(tx *gorm.DB) error {
	u.Updated = utils.DateNow()
	return nil
}
