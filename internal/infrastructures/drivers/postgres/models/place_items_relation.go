package models

import "gorm.io/gorm"

// PlaceItems represents the relationship between a Place and its Items,
// including the count of items in the relationship.
type PlaceItems struct {
	PlaceID   uint `gorm:"not null"`  // Foreign Key to Place
	ItemID    uint `gorm:"not null"`  // Foreign Key to Item
	NbOfItems uint `gorm:"default:0"` // Count of items in this relationship
}

type PlaceItemsDAO struct {
	DBClient *gorm.DB
}
