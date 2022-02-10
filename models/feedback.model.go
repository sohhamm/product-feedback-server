package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string
	Category    string
	Upvotes     uint
	Status      string
	Description string
	// Comments Comment
}
