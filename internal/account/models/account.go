package models

type AccountRequest struct {
	AccountID string `params:"account_id" validate:"required,uuid"`
}

type AccountResponse struct {
	AccountID string `json:"account_id"`
	Type     string `json:"type"`
	Currency string `json:"currency"`
	Issuer  string `json:"issuer"`
}
