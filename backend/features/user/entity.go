package user

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserCore struct {
	UserID    string
	Username  string
	Email     string
	Password  string
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
	Login() gin.HandlerFunc
}

type UseCase interface {
	Register(request UserCore) (UserCore, error)
	Login(request UserCore) (UserCore, string, error)
}

type Repository interface {
	Register(request UserCore) (UserCore, error)
	Login(request UserCore) (UserCore, string, error)
}
