package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_REGISTER_USER   = "failed create user"
	MESSAGE_FAILED_PROSES_REQUEST  = "failed proses request"
	MESSAGE_FAILED_TOKEN_NOT_FOUND = "token not found"
	MESSAGE_FAILED_TOKEN_NOT_VALID = "token not valid"
	MESSAGE_FAILED_DENIED_ACCESS   = "denied access"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER = "success create user"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCreateUser         = errors.New("failed to create user")
)

type (
	UserRegisterRequest struct {
		Name       string `json:"name" form:"name"`
		TelpNumber string `json:"telp_number" form:"telp_number"`
		Email      string `json:"email" form:"email"`
		Password   string `json:"password" form:"password"`
	}

	UserRegisterResponse struct {
		Name       string `json:"name"`
		TelpNumber string `json:"telp_number"`
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}
)
