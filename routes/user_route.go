package routes

import (
	"github.com/Shabrinashsf/ets-backend-webpro-c/controller"
	"github.com/Shabrinashsf/ets-backend-webpro-c/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/user")
	{
		routes.POST("/register", userController.Register)
		routes.POST("/login", userController.Login)
	}
}
