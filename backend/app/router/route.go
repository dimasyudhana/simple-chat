package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	uc "github.com/dimasyudhana/simple-chat/features/user/controller"
	ur "github.com/dimasyudhana/simple-chat/features/user/repository"
	uu "github.com/dimasyudhana/simple-chat/features/user/usecase"
)

func InitRouter(db *gorm.DB, r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	initUserRouter(db, r)
}

func initUserRouter(db *gorm.DB, r *gin.Engine) {
	userRepository := ur.New(db)
	userUsecase := uu.New(userRepository)
	userController := uc.New(userUsecase)

	r.POST("/register", userController.Register())
	// r.POST("/login", userController.Login())
}
