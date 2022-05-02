package models

type Comment struct {
	CreatedAt  int  `json:"-"`
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Content    string
	FeedbackID uint `json:"-"`
	UserID     int  `json:"user_id" gorm:""`
}

type ReqComment struct {
}
