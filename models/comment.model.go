package models

type Comment struct {
	CreatedAt  int  `json:"-"`
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Content    string
	FeedbackID uint `json:"-"`
}
