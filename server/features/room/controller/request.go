package controller

type RegisterRequest struct {
	RoomID   string `json:"room_id" form:"room_id"`
	RoomName string `json:"room_name" form:"room_name"`
}
