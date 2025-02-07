package models

type Banner struct {
	BannerID    string  `gorm:"primaryKey" json:"banner_id"`
	UserID      *string `gorm:"column:user_id" json:"user_id"`
	Title       *string `gorm:"column:title" json:"title"`
	Description *string `gorm:"column:description" json:"description"`
	Image       *string `gorm:"column:image" json:"image"`
	DummyCol11  *string `gorm:"column:dummy_col_11" json:"dummy_col_11"`
}
