package entity

import (
	"github.com/google/uuid"
)

type RoomType struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name  string    `json:"name"` // RoomType name, ex: Standard, Deluxe, Suite, Presidential
	Price int       `json:"price"`

	Timestamp
}
