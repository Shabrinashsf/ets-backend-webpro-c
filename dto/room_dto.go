package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_ADD_ROOM    = "Failed to add room"
	MESSAGE_FAILED_UPDATE_ROOM = "Failed to update room"

	// Success
	MESSAGE_SUCCESS_ADD_ROOM    = "Success to add room"
	MESSAGE_SUCCESS_UPDATE_ROOM = "Success to update room"
)

var (
	ErrRoomAlreadyExists = errors.New("room already exists")
	ErrParsedUUID        = errors.New("failed to parse uuid")
	ErrCreateRoom        = errors.New("failed to create room")
	ErrRoomNotFound      = errors.New("room not found")
	ErrUpdateRoom        = errors.New("failed to update room")
)

type (
	BookedRoomRequest struct {
	}

	AddRoomRequest struct {
		Number int    `json:"number" form:"number"`
		TypeID string `json:"type_id" form:"type_id"`
	}

	AddRoomResponse struct {
		Number int    `json:"number" form:"number"`
		TypeID string `json:"type_id" form:"type_id"`
		Status string `json:"status" form:"status"`
	}

	UpdateRoomRequest struct {
		TypeID string `json:"type_id" form:"type_id"`
		Status string `json:"status" form:"status"`
	}
)
