package message

import (
	"time"

	"gorm.io/gorm"
)

type MessageCore struct {
	MessageID string
	UserID    string
	RoomID    string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
