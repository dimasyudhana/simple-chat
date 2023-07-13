package member

import (
	"time"

	"gorm.io/gorm"
)

type MemberCore struct {
	MemberID  string
	UserID    string
	RoomID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
