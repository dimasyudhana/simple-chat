package database

import (
	member "github.com/dimasyudhana/simple-chat/features/member/repository"
	message "github.com/dimasyudhana/simple-chat/features/message/repository"
	room "github.com/dimasyudhana/simple-chat/features/room/repository"
	user "github.com/dimasyudhana/simple-chat/features/user/repository"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&user.User{},
		&room.Room{},
		&member.Member{},
		&message.Message{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return err
}
