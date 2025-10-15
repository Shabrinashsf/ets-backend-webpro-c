package entity

import "github.com/google/uuid"

type Room struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RoomTypeID uuid.UUID `json:"room_type_id"`
	Number     int       `json:"room_number"` // ex: 101, 201, 301, etc
	Status     string    `json:"status"`      // ex: booked, available

	RoomType *RoomType `gorm:"foreignKey:RoomTypeID"`

	Timestamp
}
