package repository

import (
	"time"

	member "github.com/dimasyudhana/simple-chat/features/member/repository"
	message "github.com/dimasyudhana/simple-chat/features/message/repository"
	"gorm.io/gorm"
)

type User struct {
	UserID    string            `gorm:"primaryKey;type:varchar(45)"`
	Fullname  string            `gorm:"type:varchar(225);not null"`
	Email     string            `gorm:"type:varchar(225);not null;unique"`
	Password  string            `gorm:"type:text;not null"`
	CreatedAt time.Time         `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time         `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt    `gorm:"index"`
	Members   []member.Member   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Messages  []message.Message `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
