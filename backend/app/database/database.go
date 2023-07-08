package database

import (
	"fmt"

	"github.com/dimasyudhana/simple-chat/app/config"
	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = middlewares.Log()

func InitDatabase(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHOST, config.DBUSER, config.DBPASSWORD, config.DBNAME, config.DBPORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	log.Info("success connected to database")

	return db
}
