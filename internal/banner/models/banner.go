package models

// BannerRequest represents the request to get a banner by userID
type BannerRequest struct {
	UserID string  `params:"user_id" validate:"required,uuid"`
}

// BannerResponse represents the response that contains banner details
type BannerResponse struct {
	Title      string `json:"title"`
	Description string `json:"description"`
	Image      string `json:"image"`
}
