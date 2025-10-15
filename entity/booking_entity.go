package entity

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	RoomID     uuid.UUID `json:"room_id"`
	TotalPrice int       `json:"total_price"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`

	User *User `gorm:"foreignKey:UserID"`
	Room *Room `gorm:"foreignKey:RoomID"`
}
