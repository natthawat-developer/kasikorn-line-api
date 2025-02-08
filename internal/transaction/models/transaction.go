package models

type GetTransactionByUserIDRequest struct {
	UserID string `params:"user_id" validate:"required,uuid"`
}

type GetTransactionByUserIDResponse struct {
	TransactionIDs []string `json:"transaction_ids"`
}

type GetTransactionDetailRequest struct {
	TransactionID string `params:"transaction_id" validate:"required,uuid"`
}

type GetTransactionDetailResponse struct {
	Name   *string `json:"name,omitempty"`
	Image  *string `json:"image,omitempty"`
	IsBank *bool   `json:"is_bank,omitempty"`
}
