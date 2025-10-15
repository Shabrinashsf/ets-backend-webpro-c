package repository

import (
	"context"

	"github.com/Shabrinashsf/ets-backend-webpro-c/entity"
	"gorm.io/gorm"
)

type (
	RoomRepository interface {
		CheckRoom(ctx context.Context, tx *gorm.DB, number int) (entity.Room, bool, error)
		AddRoom(ctx context.Context, tx *gorm.DB, room entity.Room) (entity.Room, error)
	}

	roomRepository struct {
		db *gorm.DB
	}
)

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) CheckRoom(ctx context.Context, tx *gorm.DB, number int) (entity.Room, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var room entity.Room
	if err := tx.WithContext(ctx).Where("number = ?", number).Take(&room).Error; err != nil {
		return entity.Room{}, false, err
	}

	return room, true, nil
}

func (r *roomRepository) AddRoom(ctx context.Context, tx *gorm.DB, room entity.Room) (entity.Room, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&room).Error; err != nil {
		return entity.Room{}, err
	}

	return room, nil
}
