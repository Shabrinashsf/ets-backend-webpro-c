package controller

import (
	"net/http"
	"strconv"

	"github.com/Shabrinashsf/ets-backend-webpro-c/dto"
	"github.com/Shabrinashsf/ets-backend-webpro-c/service"
	"github.com/Shabrinashsf/ets-backend-webpro-c/utils/response"
	"github.com/gin-gonic/gin"
)

type (
	RoomController interface {
		AddRoom(ctx *gin.Context)
		UpdateRoom(ctx *gin.Context)
		DeleteRoom(ctx *gin.Context)
		GetAllRoom(ctx *gin.Context)
		BookingRoom(ctx *gin.Context)
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

func (c *roomController) UpdateRoom(ctx *gin.Context) {
	var room dto.UpdateRoomRequest
	if err := ctx.ShouldBind(&room); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_ROOM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	idparam := ctx.Param("id")
	result, err := c.roomService.UpdateRoom(ctx.Request.Context(), room, idparam)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_ROOM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_ROOM, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *roomController) DeleteRoom(ctx *gin.Context) {
	idparam := ctx.Param("id")
	result, err := c.roomService.DeleteRoom(ctx.Request.Context(), idparam)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_ROOM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_ROOM, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *roomController) GetAllRoom(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	result, err := c.roomService.GetAllRoom(ctx.Request.Context(), page, limit)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ROOMS, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ROOMS, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *roomController) BookingRoom(ctx *gin.Context) {
	var room dto.BookingRoomRequest
	userID := ctx.MustGet("user_id").(string)

	if err := ctx.ShouldBind(&room); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_ADD_ROOM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.roomService.BookingRoom(ctx.Request.Context(), room, userID)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_BOOKING_ROOM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(dto.MESSAGE_SUCCESS_BOOKING_ROOM, result)
	ctx.JSON(http.StatusOK, res)
}
