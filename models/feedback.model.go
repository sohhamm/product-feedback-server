package models

import (
	"github.com/google/uuid"
)

type Feedback struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid"`
	Title       string    `json:"title" gorm:"not null"`
	Category    string    `json:"category" gorm:"not null"`
	Upvotes     uint      `json:"upvotes" gorm:"default:0"`
	Status      string    `json:"status" gorm:"default:suggestion"`
	Description string    `json:"description" gorm:"not null"`
	// Comments Comment
}
