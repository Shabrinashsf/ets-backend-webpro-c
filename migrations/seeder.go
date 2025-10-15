package migrations

import (
	"github.com/Shabrinashsf/ets-backend-webpro-c/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListRoomTypeSeeder(db); err != nil {
		return err
	}

	return nil
}
