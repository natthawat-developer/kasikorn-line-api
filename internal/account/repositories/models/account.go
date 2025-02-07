package models

import (
	"time"
)

// AccountBalance represents the account_balances table
type AccountBalance struct {
	AccountID string  `gorm:"primaryKey;column:account_id"`
	UserID    *string `gorm:"column:user_id"`
	Amount    *float64 `gorm:"column:amount"`
	DummyCol4 *string `gorm:"column:dummy_col_4"`
}

// AccountDetail represents the account_details table
type AccountDetail struct {
	AccountID     string  `gorm:"primaryKey;column:account_id"`
	UserID        *string `gorm:"column:user_id"`
	Color         *string `gorm:"column:color"`
	IsMainAccount *bool   `gorm:"column:is_main_account"`
	Progress      *int    `gorm:"column:progress"`
	DummyCol5     *string `gorm:"column:dummy_col_5"`
}

// AccountFlag represents the account_flags table
type AccountFlag struct {
	FlagID    int       `gorm:"primaryKey;autoIncrement;column:flag_id"`
	AccountID string    `gorm:"column:account_id"`
	UserID    string    `gorm:"column:user_id"`
	FlagType  string    `gorm:"column:flag_type"`
	FlagValue string    `gorm:"column:flag_value"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

// Account represents the accounts table
type Account struct {
	AccountID     string  `gorm:"primaryKey;column:account_id"`
	UserID        *string `gorm:"column:user_id"`
	Type          *string `gorm:"column:type"`
	Currency      *string `gorm:"column:currency"`
	AccountNumber *string `gorm:"column:account_number"`
	Issuer        *string `gorm:"column:issuer"`
	DummyCol3     *string `gorm:"column:dummy_col_3"`
}
