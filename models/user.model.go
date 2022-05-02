package models

type User struct {
	Name     string
	Username string `gorm:"unique"`
	GoogleID string
}
