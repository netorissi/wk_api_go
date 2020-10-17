package sqlstore

import (
	"github.com/netorissi/wk_api_go/entities"
	"github.com/netorissi/wk_api_go/utils"
)

func (sql *SqlSupplier) Migrate() {
	db := sql.GetConn()
	if db == nil {
		return
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&entities.Session{},
		&entities.Avatar{},
		&entities.User{},
		&entities.UserInformation{},
	)

	db.Create(&entities.Avatar{
		Name:    "Default",
		Type:    1,
		URL:     "https://img.icons8.com/cotton/2x/worldwide-location.png",
		Created: utils.DateNow(),
	})
}
