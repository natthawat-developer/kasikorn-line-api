package models

type UserRequest struct {
	UserID string `params:"user_id" validate:"required,uuid"`
}

type UserResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}
