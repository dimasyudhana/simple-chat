package database

import (
	message "github.com/dimasyudhana/simple-chat/features/message/repository"
	room "github.com/dimasyudhana/simple-chat/features/room/repository"
	user "github.com/dimasyudhana/simple-chat/features/user/repository"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	err := db.SetupJoinTable(&user.User{}, "Rooms", &user.Members{})
	if err != nil {
		log.Sugar().Error("setup err ", err)
		return err
	}
	err = db.AutoMigrate(
		&user.User{},
		&room.Room{},
		&message.Message{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("success migrated to database")

	return err
}
