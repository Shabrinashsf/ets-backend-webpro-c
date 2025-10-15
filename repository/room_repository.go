package repository

import (
	"context"

	"github.com/Shabrinashsf/ets-backend-webpro-c/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	RoomRepository interface {
		CheckRoom(ctx context.Context, tx *gorm.DB, number int) (entity.Room, bool, error)
		AddRoom(ctx context.Context, tx *gorm.DB, room entity.Room) (entity.Room, error)
		GetRoomByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Room, error)
		UpdateRoom(ctx context.Context, tx *gorm.DB, room entity.Room) (entity.Room, error)
		DeleteRoom(ctx context.Context, tx *gorm.DB, room entity.Room) error
		GetRoomTypeByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.RoomType, error)
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

func (r *roomRepository) GetRoomByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Room, error) {
	if tx == nil {
		tx = r.db
	}

	var room entity.Room
	if err := tx.WithContext(ctx).Where("id = ?", id).Take(&room).Error; err != nil {
		return entity.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) UpdateRoom(ctx context.Context, tx *gorm.DB, room entity.Room) (entity.Room, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Model(&room).Updates(room).Error; err != nil {
		return entity.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) DeleteRoom(ctx context.Context, tx *gorm.DB, room entity.Room) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&room).Error; err != nil {
		return err
	}

	return nil
}

func (r *roomRepository) GetRoomTypeByID(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.RoomType, error) {
	if tx == nil {
		tx = r.db
	}

	var roomType entity.RoomType
	if err := tx.WithContext(ctx).Where("id = ?", id).Take(&roomType).Error; err != nil {
		return entity.RoomType{}, err
	}

	return roomType, nil
}
