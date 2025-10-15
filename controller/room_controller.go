package controller

import (
	"net/http"

	"github.com/Shabrinashsf/ets-backend-webpro-c/dto"
	"github.com/Shabrinashsf/ets-backend-webpro-c/service"
	"github.com/Shabrinashsf/ets-backend-webpro-c/utils/response"
	"github.com/gin-gonic/gin"
)

type (
	RoomController interface {
		AddRoom(ctx *gin.Context)
	}

	roomController struct {
		roomService service.RoomService
	}
)

func NewRoomController(roomService service.RoomService) RoomController {
	return &roomController{
		roomService: roomService,
	}
}

func (c *roomController) AddRoom(ctx *gin.Context) {
	var room dto.AddRoomRequest
	if err := ctx.ShouldBind(&room); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_ADD_ROOM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.roomService.AddRoom(ctx.Request.Context(), room)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_ADD_ROOM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(dto.MESSAGE_SUCCESS_ADD_ROOM, result)
	ctx.JSON(http.StatusOK, res)
}
