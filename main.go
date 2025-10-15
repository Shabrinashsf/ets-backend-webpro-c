package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Shabrinashsf/ets-backend-webpro-c.git/cmd"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/config"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/controller"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/middleware"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/repository"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/routes"
	"github.com/Shabrinashsf/ets-backend-webpro-c.git/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Commands(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(db)

		// Service
		userService service.UserService = service.NewUserService(userRepository, jwtService)

		// Controller
		userController controller.UserController = controller.NewUserController(userService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Route Not Found",
		})
	})

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// routes
	routes.User(server, userController, jwtService)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
