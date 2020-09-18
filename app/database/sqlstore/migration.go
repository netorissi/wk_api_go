package sqlstore

import "github.com/netorissi/wk_api_go/entities"

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
}
