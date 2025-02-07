package models

import "time"

type GetAccountDetailRequest struct {
	AccountID string `params:"account_id" validate:"required,uuid"`
}

type GetAccountDetailResponse struct {
	Type          *string       `json:"type"`
	Currency      *string       `json:"currency"`
	AccountNumber *string       `json:"account_number"`
	Issuer        *string       `json:"issuer"`
	Amount        *float64      `json:"amount"`
	Color         *string       `json:"color"`
	IsMainAccount *bool         `json:"is_main_account"`
	Progress      *int          `json:"progress"`
	Flags         []AccountFlag `json:"flags"`
}

type AccountFlag struct {
	FlagType  string    `json:"flag_type"`
	FlagValue string    `json:"flag_value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountDetailResponse struct {
	Color         string `json:"color"`
	IsMainAccount string `json:"is_main_account"`
	Progress      string `json:"progress"`
}

type GetAccountByUserIDRequest struct {
	UserID string `params:"user_id" validate:"required,uuid"`
}

type GetAccountByUserIDResponse struct {
	AccountIDs []string `json:"account_ids"`
}
