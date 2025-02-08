package models

type Transaction struct {
	TransactionID string  `gorm:"primaryKey;column:transaction_id"`
	UserID        *string `gorm:"column:user_id"`
	Name          *string `gorm:"column:name"`
	Image         *string `gorm:"column:image"`
	IsBank        *bool   `gorm:"column:is_bank"`
	DummyCol6     *string `gorm:"column:dummy_col_6"`
}
