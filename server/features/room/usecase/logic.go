package usecase

import (
	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/dimasyudhana/simple-chat/features/room"
)

var log = middlewares.Log()

type Service struct {
	query room.Repository
}

func New(ud room.Repository) room.UseCase {
	return &Service{
		query: ud,
	}
}
