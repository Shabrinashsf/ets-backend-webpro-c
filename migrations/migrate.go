package migrations

import (
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		return err
	}

	return nil
}
