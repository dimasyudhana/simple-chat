package room

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RoomCore struct {
	RoomID    string
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Members   []MemberCore
	Messages  []MessageCore
}

type MemberCore struct {
	MemberID  string
	UserID    string
	RoomID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageCore struct {
	MessageID string
	UserID    string
	RoomID    string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Controller interface {
	Register() gin.HandlerFunc
	Join() gin.HandlerFunc
	GetRooms() gin.HandlerFunc
	GetMembers() gin.HandlerFunc
}

type UseCase interface {
}

type Repository interface {
}
