package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Shabrinashsf/ets-backend-webpro-c/cmd"
	"github.com/Shabrinashsf/ets-backend-webpro-c/config"
	"github.com/Shabrinashsf/ets-backend-webpro-c/controller"
	"github.com/Shabrinashsf/ets-backend-webpro-c/middleware"
	"github.com/Shabrinashsf/ets-backend-webpro-c/repository"
	"github.com/Shabrinashsf/ets-backend-webpro-c/routes"
	"github.com/Shabrinashsf/ets-backend-webpro-c/service"
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
		roomRepository repository.RoomRepository = repository.NewRoomRepository(db)

		// Service
		userService service.UserService = service.NewUserService(userRepository, jwtService)
		roomService service.RoomService = service.NewRoomService(roomRepository, userRepository)

		// Controller
		userController controller.UserController = controller.NewUserController(userService)
		roomController controller.RoomController = controller.NewRoomController(roomService)
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
	routes.Room(server, roomController, jwtService)

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
