package routes

import (
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/controller"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/user")
	{
		routes.POST("/register", userController.Register)
	}
}
