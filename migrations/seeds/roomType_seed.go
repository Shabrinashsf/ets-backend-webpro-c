package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Shabrinashsf/ets-backend-webpro-c/entity"
	"gorm.io/gorm"
)

func ListRoomTypeSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/roomType.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listRoomType []entity.RoomType
	if err := json.Unmarshal(jsonData, &listRoomType); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.RoomType{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.RoomType{}); err != nil {
			return err
		}
	}

	for _, data := range listRoomType {
		var roomType entity.RoomType
		err := db.Where(&entity.RoomType{Name: data.Name}).First(&roomType).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&roomType, "name = ?", data.Name).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
