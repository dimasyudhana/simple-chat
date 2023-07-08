package main

import (
	"github.com/dimasyudhana/simple-chat/app/config"
	"github.com/dimasyudhana/simple-chat/app/database"
	"github.com/dimasyudhana/simple-chat/app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	cfg := config.InitConfig()
	db := database.InitDatabase(cfg)
	database.InitMigration(db)
	router.InitRouter(db, r)
	r.Run(":8080")
}
