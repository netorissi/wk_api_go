package entities

import (
	"strings"

	"github.com/netorissi/wk_api_go/utils"
	"gorm.io/gorm"
)

type Avatar struct {
	ID      int    `gorm:"type:int;auto_increment;primary_key;" json:"id"`
	Name    string `gorm:"type:varchar(200);unique;default:''" json:"name"`
	Type    int    `gorm:"type:int(1);default:0" json:"type"`
	URL     string `gorm:"type:varchar(255);default:'';not null" json:"url"`
	Updated int64  `gorm:"autoUpdateTime;default:0" json:"updated"`
	Created int64  `gorm:"autoCreateTime;default:0" json:"created"`
}

func (a *Avatar) BeforeCreate(tx *gorm.DB) error {
	a.URL = strings.ToLower(a.URL)
	a.Created = utils.DateNow()
	return nil
}
