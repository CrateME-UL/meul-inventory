// Package models provides the templates for the models.
package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ItemID     uint      `gorm:"primaryKey"`
	PlaceID    uint      `gorm:"foreignKey:PlaceID"`
	ItemNumber uuid.UUID `gorm:"type:uuid;unique;not null"`
	Name       string    `gorm:"not null;"`

	DB *gorm.DB
}
