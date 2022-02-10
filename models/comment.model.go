package models

type Comment struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Content string
}
