package repository

import (
	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/dimasyudhana/simple-chat/features/room"
	"gorm.io/gorm"
)

var log = middlewares.Log()

type Query struct {
	db *gorm.DB
}

func New(db *gorm.DB) room.Repository {
	return &Query{
		db: db,
	}
}
