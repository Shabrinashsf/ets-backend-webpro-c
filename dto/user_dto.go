package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_REGISTER_USER   = "failed create user"
	MESSAGE_FAILED_PROSES_REQUEST  = "failed proses request"
	MESSAGE_FAILED_TOKEN_NOT_FOUND = "token not found"
	MESSAGE_FAILED_TOKEN_NOT_VALID = "token not valid"
	MESSAGE_FAILED_DENIED_ACCESS   = "denied access"
	MESSAGE_FAILED_LOGIN_USER      = "failed login user"
	MESSAGE_FAILED_GET_USER        = "failed get user"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER = "success create user"
	MESSAGE_SUCCESS_LOGIN_USER    = "success login user"
	MESSAGE_SUCCESS_GET_USER      = "success get user"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCreateUser         = errors.New("failed to create user")
	ErrEmailNotFound      = errors.New("email not found")
	ErrAccountNotVerified = errors.New("account not verified")
	ErrPasswordNotMatch   = errors.New("password not match")
	ErrRoleNotAllowed     = errors.New("role not allowed")
	ErrGetUserById        = errors.New("failed get user by id")
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

	UserLoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}

	GetMeResponse struct {
		Name       string `json:"name"`
		TelpNumber string `json:"telp_number"`
		Email      string `json:"email"`
	}
)
