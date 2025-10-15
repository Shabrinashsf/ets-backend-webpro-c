package routes

import (
	"github.com/Shabrinashsf/ets-backend-webpro-c/constants"
	"github.com/Shabrinashsf/ets-backend-webpro-c/controller"
	"github.com/Shabrinashsf/ets-backend-webpro-c/middleware"
	"github.com/Shabrinashsf/ets-backend-webpro-c/service"
	"github.com/gin-gonic/gin"
)

func Room(route *gin.Engine, roomController controller.RoomController, jwtService service.JWTService) {
	routes := route.Group("/room")
	{
		routes.POST("/add", middleware.Authenticate(jwtService), middleware.OnlyAllow(string(constants.ENUM_ROLE_ADMIN)), roomController.AddRoom)
		routes.PUT("/:id", middleware.Authenticate(jwtService), middleware.OnlyAllow(string(constants.ENUM_ROLE_ADMIN)), roomController.UpdateRoom)
		routes.DELETE("/:id", middleware.Authenticate(jwtService), middleware.OnlyAllow(string(constants.ENUM_ROLE_ADMIN)), roomController.DeleteRoom)
		routes.GET("/", roomController.GetAllRoom)
	}
}
