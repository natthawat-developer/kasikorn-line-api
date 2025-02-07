package models

import "time"

type GetAccountDetailRequest struct {
	AccountID string `params:"account_id" validate:"required,uuid"`
}

type GetAccountDetailResponse struct {
	Type          *string       `json:"type"`
	Currency      *string       `json:"currency"`
	Issuer        *string       `json:"issuer"`
	Balance       *float64      `json:"balance"`
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
