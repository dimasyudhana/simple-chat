package repository

import (
	"time"

	room "github.com/dimasyudhana/simple-chat/features/room/repository"
	"github.com/dimasyudhana/simple-chat/features/user"
	"gorm.io/gorm"
)

type User struct {
	UserID    string         `gorm:"primaryKey;type:varchar(45)"`
	Username  string         `gorm:"type:varchar(225);not null"`
	Email     string         `gorm:"type:varchar(225);not null;unique"`
	Password  string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Messages  []room.Room    `gorm:"many2many:Messages;foreignKey:UserID;joinForeignKey:UserID"`
	Members   []room.Room    `gorm:"many2many:Members;foreignKey:UserID;joinForeignKey:UserID"`
}

type Members struct {
	MemberID   string         `gorm:"primaryKey;type:varchar(45)"`
	UserID     string         `gorm:"foreignKey;type:varchar(45)"`
	RoomRoomID string         `gorm:"foreignKey;type:varchar(45)"`
	CreatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Messages struct {
	MessageID  string         `gorm:"primaryKey;type:varchar(45)"`
	UserID     string         `gorm:"foreignKey;type:varchar(45)"`
	RoomRoomID string         `gorm:"foreignKey;type:varchar(45)"`
	Message    string         `gorm:"type:varchar(225);not null"`
	CreatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func modelToCore(u User) user.UserCore {
	return user.UserCore{
		UserID:    u.UserID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func coreToModel(u user.UserCore) User {
	return User{
		UserID:    u.UserID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
