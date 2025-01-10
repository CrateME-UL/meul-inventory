package fixtures

import (
	"meul/inventory/internal/infrastructures/drivers/postgres/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

// Option defines a function type for modifying an Item
type Option func(*models.Item)

// NewItemFixture creates an Item with default random values and applies given options
func NewItemFixture(options ...Option) models.Item {
	rand.Seed(uint64(time.Now().UnixNano()))

	// Default values
	item := models.Item{
		ItemID:     uint(rand.Intn(1000)),    // Random ItemID
		ItemNumber: uuid.New(),               // Random UUID
		Name:       generateRandomString(10), // Random 10-character string
	}

	// Apply options
	for _, opt := range options {
		opt(&item)
	}

	return item
}

// WithItemID sets a specific ItemID
func WithItemID(itemID uint) Option {
	return func(i *models.Item) {
		i.ItemID = itemID
	}
}

// WithItemNumber sets a specific ItemNumber
func WithItemNumber(itemNumber uuid.UUID) Option {
	return func(i *models.Item) {
		i.ItemNumber = itemNumber
	}
}

// WithName sets a specific Name
func WithName(name string) Option {
	return func(i *models.Item) {
		i.Name = name
	}
}
