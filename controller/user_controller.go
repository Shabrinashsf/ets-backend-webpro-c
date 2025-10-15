package controller

import (
	"net/http"

	"github.com/Shabrinashsf/ets-backend-webpro-c.git/dto"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/service"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/utils/response"
	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user dto.UserRegisterRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.Register(ctx.Request.Context(), user)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}
