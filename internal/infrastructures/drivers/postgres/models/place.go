package models

import (
	"github.com/google/uuid"
)

// Place represents a template for a storage location.
type Place struct {
	PlaceID     uint      `gorm:"primaryKey"`                // Primary Key
	PlaceNumber uuid.UUID `gorm:"type:uuid;unique;not null"` // Unique Identifier
	Name        string    `gorm:"unique;size:30;not null;"`  // Name of the Place
}
