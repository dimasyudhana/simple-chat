package controller

import "github.com/dimasyudhana/simple-chat/features/user"

type RegisterResponse struct {
	UserID string `json:"user_id,omitempty"`
	Email  string `json:"email,omitempty"`
}

func CoreToResponse(data interface{}) RegisterResponse {
	res := RegisterResponse{}
	switch v := data.(type) {
	case user.UserCore:
		res.UserID = v.UserID
		res.Email = v.Email
	default:
		return RegisterResponse{}
	}

	return res
}
