package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	rc "github.com/dimasyudhana/simple-chat/features/room/controller"
	rr "github.com/dimasyudhana/simple-chat/features/room/repository"
	ru "github.com/dimasyudhana/simple-chat/features/room/usecase"
	uc "github.com/dimasyudhana/simple-chat/features/user/controller"
	ur "github.com/dimasyudhana/simple-chat/features/user/repository"
	uu "github.com/dimasyudhana/simple-chat/features/user/usecase"
	"github.com/dimasyudhana/simple-chat/utils/websockets"
)

func InitRouter(db *gorm.DB, r *gin.Engine, hub *websockets.Hub) {
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
	initRoomRouter(db, r, hub)
}

func initUserRouter(db *gorm.DB, r *gin.Engine) {
	userRepository := ur.New(db)
	userUsecase := uu.New(userRepository)
	userController := uc.New(userUsecase)

	r.POST("/signup", userController.Register())
	r.POST("/login", userController.Login())
	r.GET("/logout", userController.Logout())
}

func initRoomRouter(db *gorm.DB, r *gin.Engine, hub *websockets.Hub) {
	roomRepository := rr.New(db)
	roomUsecase := ru.New(roomRepository)
	roomController := rc.New(roomUsecase, hub)

	r.POST("/rooms/register", roomController.Register())
	r.GET("/rooms/:id/join", roomController.Join())
}
