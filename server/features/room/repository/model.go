package repository

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	RoomID    string         `gorm:"primaryKey;type:varchar(45)"`
	RoomName  string         `gorm:"type:varchar(225);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Members   []Members      `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Messages  []Messages     `gorm:"foreignKey:RoomID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Members struct {
	MemberID  string         `gorm:"primaryKey;type:varchar(45)"`
	UserID    string         `gorm:"foreignKey;type:varchar(45)"`
	RoomID    string         `gorm:"foreignKey;type:varchar(45)"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Messages struct {
	MessageID string         `gorm:"primaryKey;type:varchar(45)"`
	UserID    string         `gorm:"foreignKey;type:varchar(45)"`
	RoomID    string         `gorm:"foreignKey;type:varchar(45)"`
	Message   string         `gorm:"type:varchar(225);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
