package entities

type UserInformation struct {
	ID               int    `gorm:"type:int;auto_increment; primary_key;" json:"id"`
	UserID           int    `gorm:"type:int;not null"`
	User             User   `gorm:"foreignKey:UserID;" json:"user"`
	EmailConfirm     string `gorm:"int(1);default:0" json:"email_confirm"`
	EmailUpdated     int64  `gorm:"autoUpdateTime;default:0" json:"email_updated"`
	PasswordUpdated  int64  `gorm:"autoUpdateTime;default:0" json:"password_updated"`
	CellphoneUpdated int64  `gorm:"autoUpdateTime;default:0" json:"cellphone_updated"`
	Updated          int64  `gorm:"autoUpdateTime;default:0" json:"updated"`
	Created          int64  `gorm:"autoCreateTime;default:0" json:"created"`
}
