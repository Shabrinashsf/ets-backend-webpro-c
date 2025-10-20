package dto

import (
	"errors"
	"time"

	"github.com/Shabrinashsf/ets-backend-webpro-c/utils/pagination"
)

const (
	// Failed
	MESSAGE_FAILED_ADD_ROOM     = "Failed to add room"
	MESSAGE_FAILED_UPDATE_ROOM  = "Failed to update room"
	MESSAGE_FAILED_DELETE_ROOM  = "Failed to delete room"
	MESSAGE_FAILED_GET_ROOMS    = "Failed to get rooms"
	MESSAGE_FAILED_BOOKING_ROOM = "Failed to book room"

	// Success
	MESSAGE_SUCCESS_ADD_ROOM     = "Success to add room"
	MESSAGE_SUCCESS_UPDATE_ROOM  = "Success to update room"
	MESSAGE_SUCCESS_DELETE_ROOM  = "Success to delete room"
	MESSAGE_SUCCESS_GET_ROOMS    = "Success to get rooms"
	MESSAGE_SUCCESS_BOOKING_ROOM = "Success to book room"
)

var (
	ErrRoomAlreadyExists    = errors.New("room already exists")
	ErrParsedUUID           = errors.New("failed to parse uuid")
	ErrCreateRoom           = errors.New("failed to create room")
	ErrRoomNotFound         = errors.New("room not found")
	ErrUpdateRoom           = errors.New("failed to update room")
	ErrDeleteRoom           = errors.New("failed to delete room")
	ErrRoomTypeNotFound     = errors.New("room type not found")
	ErrRoomNotAvailable     = errors.New("room is not available for booking")
	ErrInvalidTimeFormat    = errors.New("invalid time format")
	ErrCheckOutAfterCheckIn = errors.New("check-out date must be after check-in date")
	ErrBookedRoom           = errors.New("failed to booked a room")
)

type (
	BookingRoomRequest struct {
		RoomID   string `json:"room_id" form:"room_id"`
		Name     string `json:"name" form:"name"`
		CheckIn  string `json:"check_in" form:"check_in"`
		CheckOut string `json:"check_out" form:"check_out"`
	}

	BookingRoomResponse struct {
		RoomNumber int    `json:"room_number"`
		Name       string `json:"name"`
		CheckIn    string `json:"check_in"`
		CheckOut   string `json:"check_out"`
		Price      int    `json:"price"`
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

	GetRoomResponse struct {
		Number       int    `json:"number"`
		RoomTypeName string `json:"room_type_name"`
		Status       string `json:"status"`
	}

	PaginatedRoomsResponse struct {
		Data       []GetRoomResponse             `json:"data"`
		Pagination pagination.PaginationResponse `json:"pagination"`
	}
)
