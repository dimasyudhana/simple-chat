package main

import (
	"github.com/dimasyudhana/simple-chat/app/config"
	"github.com/dimasyudhana/simple-chat/app/database"
	"github.com/dimasyudhana/simple-chat/app/router"
	websockets "github.com/dimasyudhana/simple-chat/utils/websockets"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg := config.InitConfig()
	db := database.InitDatabase(cfg)
	database.InitMigration(db)
	router.InitRouter(db, r)
	go websockets.New().Run()
	r.Run(":8080")
}
