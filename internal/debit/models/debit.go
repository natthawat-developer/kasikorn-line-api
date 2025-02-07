package models

type GetDebitCardDetailsByUserIDRequest struct {
	UserID string `params:"user_id" validate:"required,uuid"`
}
type GetDebitCardDetailsByUserIDResponse struct {
	Name        *string `json:"name"`
	Color       *string `json:"color"`
	BorderColor *string `json:"border_color"`
	Issuer      *string `json:"issuer"`
	Number      *string `json:"number"`
	Status      *string `json:"status"`
}
