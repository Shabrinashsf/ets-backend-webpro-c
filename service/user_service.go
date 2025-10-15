package service

import (
	"context"

	"github.com/Shabrinashsf/ets-backend-webpro-c.git/constants"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/dto"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/entity"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/repository"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/utils/mailer"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.UserRegisterRequest) (dto.UserRegisterResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *userService) Register(ctx context.Context, req dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {
	_, flag, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.UserRegisterResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.User{
		Name:       req.Name,
		TelpNumber: req.TelpNumber,
		Email:      req.Email,
		Password:   req.Password,
		Role:       constants.ENUM_ROLE_USER,
		IsVerified: false,
	}

	userReg, err := s.userRepo.Register(ctx, nil, user)
	if err != nil {
		return dto.UserRegisterResponse{}, dto.ErrCreateUser
	}

	draftEmail, err := mailer.MakeVerificationEmail(userReg.Email)
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}

	err = mailer.SendMail(userReg.Email, draftEmail["subject"], draftEmail["body"])
	if err != nil {
		return dto.UserRegisterResponse{}, err
	}

	return dto.UserRegisterResponse{
		Name:       userReg.Name,
		TelpNumber: userReg.TelpNumber,
		Email:      userReg.Email,
		IsVerified: userReg.IsVerified,
	}, nil
}
