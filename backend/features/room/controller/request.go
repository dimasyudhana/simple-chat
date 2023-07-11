package controller

type RegisterRoomRequest struct {
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
}
