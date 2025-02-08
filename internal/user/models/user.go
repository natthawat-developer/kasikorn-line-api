package models

type UserRequest struct {
	UserID string `params:"user_id" validate:"required,uuid"`
}

type UserResponse struct {
	Name   string `json:"name"`
	Greeting string `json:"greeting"`
}
