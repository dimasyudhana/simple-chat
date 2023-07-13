package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	uc "github.com/dimasyudhana/simple-chat/features/user/controller"
	ur "github.com/dimasyudhana/simple-chat/features/user/repository"
	uu "github.com/dimasyudhana/simple-chat/features/user/usecase"
	websockets "github.com/dimasyudhana/simple-chat/utils/websocket"
)

func InitRouter(db *gorm.DB, r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	initUserRouter(db, r)
	initWebsocketRoutes(r)
}

func initUserRouter(db *gorm.DB, r *gin.Engine) {
	userRepository := ur.New(db)
	userUsecase := uu.New(userRepository)
	userController := uc.New(userUsecase)

	r.POST("/signup", userController.Register())
	r.POST("/login", userController.Login())
	r.GET("/logout", userController.Logout())
}

func initWebsocketRoutes(r *gin.Engine) {
	hub := websockets.NewHub()
	websocketHandler := websockets.NewHandler(hub)

	r.POST("/rooms/register", websocketHandler.RegisterRoom())
	r.GET("/rooms/:roomId/join", websocketHandler.JoinRoom())
}
