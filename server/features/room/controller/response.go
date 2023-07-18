package controller

type RoomResponse struct {
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
}

type MemberResponse struct {
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}
