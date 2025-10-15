package dto

import (
	"errors"
	"time"
)

const (
	// Failed
	MESSAGE_FAILED_ADD_ROOM    = "Failed to add room"
	MESSAGE_FAILED_UPDATE_ROOM = "Failed to update room"
	MESSAGE_FAILED_DELETE_ROOM = "Failed to delete room"

	// Success
	MESSAGE_SUCCESS_ADD_ROOM    = "Success to add room"
	MESSAGE_SUCCESS_UPDATE_ROOM = "Success to update room"
	MESSAGE_SUCCESS_DELETE_ROOM = "Success to delete room"
)

var (
	ErrRoomAlreadyExists = errors.New("room already exists")
	ErrParsedUUID        = errors.New("failed to parse uuid")
	ErrCreateRoom        = errors.New("failed to create room")
	ErrRoomNotFound      = errors.New("room not found")
	ErrUpdateRoom        = errors.New("failed to update room")
	ErrDeleteRoom        = errors.New("failed to delete room")
	ErrRoomTypeNotFound  = errors.New("room type not found")
)

type (
	BookedRoomRequest struct {
	}

	AddRoomRequest struct {
		Number int    `json:"number" form:"number"`
		TypeID string `json:"type_id" form:"type_id"`
	}

	AddRoomResponse struct {
		Number int    `json:"number"`
		TypeID string `json:"type_id"`
		Status string `json:"status"`
	}

	UpdateRoomRequest struct {
		TypeID string `json:"type_id" form:"type_id"`
		Status string `json:"status" form:"status"`
	}

	Timestamp struct {
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at,omitempty"`
	}

	DeleteRoomResponse struct {
		RoomTypeName string `json:"room_type_name"`
		Number       int    `json:"number"`
		Status       string `json:"status"`
		Timestamp
	}
)
