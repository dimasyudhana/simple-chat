package router

import (
	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, r *gin.Engine) {
	r = gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.TrailingSlash())
	// initUserRouter(db, r)
}

// func initUserRouter(db *gorm.DB, r *gin.Engine) {
// 	userData := ud.New(db)
// 	userService := us.New(userData)
// 	userHandler := uh.New(userService)

// 	r.POST("/register", userHandler.Register())
// 	r.POST("/login", userHandler.Login())
// }
