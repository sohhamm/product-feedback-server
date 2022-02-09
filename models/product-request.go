package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRequest struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Title    string
	Category string
	Upvotes  uint
	Status   string
	// Comments Comment
}
