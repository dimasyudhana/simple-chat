package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	uc "github.com/dimasyudhana/simple-chat/features/user/controller"
	ur "github.com/dimasyudhana/simple-chat/features/user/repository"
	uu "github.com/dimasyudhana/simple-chat/features/user/usecase"
	websockets "github.com/dimasyudhana/simple-chat/utils/websocket"
)

func InitRouter(db *gorm.DB, r *gin.Engine) {
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	initUserRouter(db, r)
}

func initUserRouter(db *gorm.DB, r *gin.Engine) {
	userRepository := ur.New(db)
	userUsecase := uu.New(userRepository)
	userController := uc.New(userUsecase)

	hub := websockets.NewHub()
	websocketHandler := websockets.NewHandler(hub)

	r.POST("/signup", userController.Register())
	r.POST("/login", userController.Login())
	r.GET("/logout", userController.Logout())

	r.POST("/registerRoom", websocketHandler.RegisterRoom())
	r.GET("/joinRoom/:roomId", websocketHandler.JoinRoom())
}
