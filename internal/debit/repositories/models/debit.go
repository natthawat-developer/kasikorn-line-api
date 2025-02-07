package models

// DebitCard represents the debit_cards table
type DebitCard struct {
	CardID    string  `gorm:"primaryKey;column:card_id"`
	UserID    *string `gorm:"column:user_id"`
	Name      *string `gorm:"column:name"`
	DummyCol7 *string `gorm:"column:dummy_col_7"`
}

// DebitCardDesign represents the debit_card_design table
type DebitCardDesign struct {
	CardID      string  `gorm:"primaryKey;column:card_id"`
	UserID      *string `gorm:"column:user_id"`
	Color       *string `gorm:"column:color"`
	BorderColor *string `gorm:"column:border_color"`
	DummyCol9   *string `gorm:"column:dummy_col_9"`
}

func (DebitCardDesign) TableName() string {
	return "debit_card_design"
}

// DebitCardDetails represents the debit_card_details table
type DebitCardDetails struct {
	CardID    string  `gorm:"primaryKey;column:card_id"`
	UserID    *string `gorm:"column:user_id"`
	Issuer    *string `gorm:"column:issuer"`
	Number    *string `gorm:"column:number"`
	DummyCol10 *string `gorm:"column:dummy_col_10"`
}

// DebitCardStatus represents the debit_card_status table
type DebitCardStatus struct {
	CardID    string  `gorm:"primaryKey;column:card_id"`
	UserID    *string `gorm:"column:user_id"`
	Status    *string `gorm:"column:status"`
	DummyCol8 *string `gorm:"column:dummy_col_8"`
}

func (DebitCardStatus) TableName() string {
	return "debit_card_status"
}