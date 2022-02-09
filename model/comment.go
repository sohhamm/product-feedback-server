package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
}
