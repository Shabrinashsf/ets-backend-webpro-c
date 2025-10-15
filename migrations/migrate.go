package migrations

import (
	"github.com/Shabrinashsf/ets-backend-webpro-c/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Room{},
		&entity.RoomType{},
	); err != nil {
		return err
	}

	return nil
}
