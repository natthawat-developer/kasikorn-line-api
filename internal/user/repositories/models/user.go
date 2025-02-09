package models

type UserGreeting struct {
	UserID    string `gorm:"primaryKey;column:user_id"`
	Greeting  string `gorm:"column:greeting"`
	DummyCol2 string `gorm:"column:dummy_col_2"`
}

type User struct {
	UserID    string `gorm:"primaryKey;column:user_id"`
	Name      string `gorm:"column:name"`
	DummyCol1 string `gorm:"column:dummy_col_1"`
}
