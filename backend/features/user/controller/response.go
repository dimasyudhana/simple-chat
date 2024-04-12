package controller

import "github.com/dimasyudhana/simple-chat/features/user"

type RegisterResponse struct {
	UserID string `json:"user_id,omitempty"`
}

type LoginResponse struct {
	UserID string `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

func Register(u user.UserCore) RegisterResponse {
	return RegisterResponse{
		UserID: u.UserID,
	}
}

func Login(u user.UserCore) LoginResponse {
	return LoginResponse{
		UserID: u.UserID,
	}
}
