package models

type Feedback struct {
	CreatedAt   int       `json:"-"`
	UpdatedAt   int       `json:"-"`
	ID          uint      `json:"id" gorm:"autoIncrement"`
	Title       string    `json:"title" gorm:"not null"`
	Category    string    `json:"category" gorm:"not null"`
	Upvotes     uint      `json:"upvotes" gorm:"default:0"`
	Status      string    `json:"status" gorm:"default:suggestion"`
	Description string    `json:"description" gorm:"not null"`
	Comments    []Comment `json:"comments"`
}

type ReqFeedback struct {
	Title       string `json:"title" gorm:"not null"`
	Category    string `json:"category" gorm:"not null"`
	Upvotes     uint   `json:"upvotes" gorm:"default:0"`
	Status      string `json:"status" gorm:"default:suggestion"`
	Description string `json:"description" gorm:"not null"`
}
