package repository

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	MemberID  string         `gorm:"primaryKey;type:varchar(45)"`
	UserID    string         `gorm:"foreignKey;type:varchar(45)"`
	RoomID    string         `gorm:"foreignKey;type:varchar(45)"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
