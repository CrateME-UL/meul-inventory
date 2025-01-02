// Package models provides the templates for the models.
package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Item represents a template for an inventory item.
type Item struct {
	ItemID     uint      `gorm:"primaryKey"`                // Primary Key
	ItemNumber uuid.UUID `gorm:"type:uuid;unique;not null"` // Unique Identifier
	Name       string    `gorm:"unique;size:30;not null;"`  // Name of the Item
}

type ItemDAO struct {
	DBClient *gorm.DB
}

func NewItemDAO(db *gorm.DB) *ItemDAO {
	return &ItemDAO{
		DBClient: db,
	}
}

func (i ItemDAO) CreateItem(item Item) error {
	if err := i.DBClient.Create(&item).Error; err != nil {
		return err
	}
	return nil
}

func (i ItemDAO) GetAllItems() ([]Item, error) {
	var items []Item
	if err := i.DBClient.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
