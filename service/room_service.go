package service

import (
	"context"

	"github.com/Shabrinashsf/ets-backend-webpro-c/dto"
	"github.com/Shabrinashsf/ets-backend-webpro-c/entity"
	"github.com/Shabrinashsf/ets-backend-webpro-c/repository"
	"github.com/Shabrinashsf/ets-backend-webpro-c/utils/pagination"
	"github.com/google/uuid"
)

type (
	RoomService interface {
		AddRoom(ctx context.Context, req dto.AddRoomRequest) (dto.AddRoomResponse, error)
		UpdateRoom(ctx context.Context, req dto.UpdateRoomRequest, idparam string) (dto.AddRoomResponse, error)
		DeleteRoom(ctx context.Context, idparam string) (dto.DeleteRoomResponse, error)
		GetAllRoom(ctx context.Context, page, limit int) (dto.PaginatedRoomsResponse, error)
	}

	roomService struct {
		roomRepo repository.RoomRepository
	}
)

func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}

func (s *roomService) AddRoom(ctx context.Context, req dto.AddRoomRequest) (dto.AddRoomResponse, error) {
	_, flag, _ := s.roomRepo.CheckRoom(ctx, nil, req.Number)
	if flag {
		return dto.AddRoomResponse{}, dto.ErrRoomAlreadyExists
	}

	parsedTypeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		return dto.AddRoomResponse{}, dto.ErrParsedUUID
	}

	room := entity.Room{
		RoomTypeID: parsedTypeID,
		Number:     req.Number,
		Status:     "available",
	}

	roomAdd, err := s.roomRepo.AddRoom(ctx, nil, room)
	if err != nil {
		return dto.AddRoomResponse{}, dto.ErrCreateRoom
	}

	return dto.AddRoomResponse{
		TypeID: roomAdd.RoomTypeID.String(),
		Number: roomAdd.Number,
		Status: roomAdd.Status,
	}, nil
}

func (s *roomService) UpdateRoom(ctx context.Context, req dto.UpdateRoomRequest, idparam string) (dto.AddRoomResponse, error) {
	room, err := s.roomRepo.GetRoomByID(ctx, nil, uuid.MustParse(idparam))
	if err != nil {
		return dto.AddRoomResponse{}, dto.ErrRoomNotFound
	}

	room.RoomTypeID = uuid.MustParse(req.TypeID)
	room.Status = req.Status

	roomUpdate, err := s.roomRepo.UpdateRoom(ctx, nil, room)
	if err != nil {
		return dto.AddRoomResponse{}, dto.ErrUpdateRoom
	}

	return dto.AddRoomResponse{
		Number: roomUpdate.Number,
		TypeID: roomUpdate.RoomTypeID.String(),
		Status: roomUpdate.Status,
	}, nil
}

func (s *roomService) DeleteRoom(ctx context.Context, idparam string) (dto.DeleteRoomResponse, error) {
	room, err := s.roomRepo.GetRoomByID(ctx, nil, uuid.MustParse(idparam))
	if err != nil {
		return dto.DeleteRoomResponse{}, dto.ErrRoomNotFound
	}

	// ngambil entitiy roomtype berdasarkan room.RoomTypeID
	roomTypeName, err := s.roomRepo.GetRoomTypeByID(ctx, nil, room.RoomTypeID)
	if err != nil {
		return dto.DeleteRoomResponse{}, dto.ErrRoomTypeNotFound
	}

	res := s.roomRepo.DeleteRoom(ctx, nil, room)
	if res != nil {
		return dto.DeleteRoomResponse{}, dto.ErrDeleteRoom
	}

	return dto.DeleteRoomResponse{
		RoomTypeName: roomTypeName.Name,
		Number:       room.Number,
		Status:       room.Status,
		Timestamp: dto.Timestamp{
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
			DeletedAt: &room.DeletedAt.Time,
		},
	}, nil
}

func (s *roomService) GetAllRoom(ctx context.Context, page, limit int) (dto.PaginatedRoomsResponse, error) {
	p := pagination.Pagination{Page: page, Limit: limit}
	offset := p.GetOffset()

	rooms, total, err := s.roomRepo.GetAllRoom(ctx, nil, offset, p.Limit)
	if err != nil {
		return dto.PaginatedRoomsResponse{}, err
	}

	var roomResponses []dto.GetRoomResponse
	for _, room := range rooms {
		roomTypeName, err := s.roomRepo.GetRoomTypeByID(ctx, nil, room.RoomTypeID)
		if err != nil {
			return dto.PaginatedRoomsResponse{}, dto.ErrRoomTypeNotFound
		}

		roomResponses = append(roomResponses, dto.GetRoomResponse{
			Number:       room.Number,
			RoomTypeName: roomTypeName.Name,
			Status:       room.Status,
		})
	}

	return dto.PaginatedRoomsResponse{
		Data:       roomResponses,
		Pagination: pagination.BuildPaginationResponse(page, limit, total),
	}, nil
}
