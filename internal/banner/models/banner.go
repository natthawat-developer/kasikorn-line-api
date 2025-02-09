package models

type BannerRequest struct {
	UserID string `params:"user_id" validate:"required,uuid"`
}

type BannerResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
