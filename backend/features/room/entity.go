package room

import "time"

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
