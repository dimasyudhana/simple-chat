package database

import (
	room "github.com/dimasyudhana/simple-chat/features/room/repository"
	user "github.com/dimasyudhana/simple-chat/features/user/repository"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	err := db.SetupJoinTable(&user.User{}, "Members", &user.Members{})
	if err != nil {
		log.Sugar().Error("setup err ", err)
		panic(err.Error())
	}

	err = db.SetupJoinTable(&user.User{}, "Messages", &user.Messages{})
	if err != nil {
		log.Sugar().Error("setup err ", err)
		panic(err.Error())
	}

	err = db.AutoMigrate(
		&user.User{},
		&room.Room{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("success migrated to database")

	return err
}
