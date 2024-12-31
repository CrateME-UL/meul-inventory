// Package models provides the templates for the models.
package models

import (
	"github.com/google/uuid"
)

// Item represents a template for an inventory item.
type Item struct {
	ItemID     uint      `gorm:"primaryKey"`                // Primary Key
	ItemNumber uuid.UUID `gorm:"type:uuid;unique;not null"` // Unique Identifier
	Name       string    `gorm:"unique;size:30;not null;"`  // Name of the Item
}
