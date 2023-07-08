package repository

import (
	"time"

	member "github.com/dimasyudhana/simple-chat/features/member/repository"
	message "github.com/dimasyudhana/simple-chat/features/message/repository"
	"gorm.io/gorm"
)

type Room struct {
	RoomID    string            `gorm:"primaryKey;type:varchar(45)"`
	RoomName  string            `gorm:"type:varchar(225);not null"`
	CreatedAt time.Time         `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time         `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt    `gorm:"index"`
	Members   []member.Member   `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Messages  []message.Message `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
