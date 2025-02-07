package models

type GetDebitCardsByUserIDRequest struct {
	UserID string `params:"user_id" validate:"required,uuid"`
}

type GetDebitCardsByUserIDResponse struct {
	CardIDs []string `json:"card_ids"`
}

type GetDebitCardDetailsByCardIDRequest struct {
	CardID string `params:"card_id" validate:"required,uuid"`
}
type GetDebitCardDetailsByCardIDResponse struct {
	Name        *string `json:"name"`
	Color       *string `json:"color"`
	BorderColor *string `json:"border_color"`
	Issuer      *string `json:"issuer"`
	Number      *string `json:"number"`
	Status      *string `json:"status"`
}
