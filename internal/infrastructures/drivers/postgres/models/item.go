// Package models provides the templates for the models.
package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

// Item represents a template for an inventory item.
type Item struct {
	ItemID     uint      `gorm:"primaryKey"`                // Primary Key
	ItemNumber uuid.UUID `gorm:"type:uuid;unique;not null"` // Unique Identifier
	Name       string    `gorm:"unique;size:30;not null;"`  // Name of the Item
	Selected   bool
}

// GenerateItemFixture generates a single Item fixture with randomized or default data
func GenerateItemFixture() Item {
	return Item{
		ItemID:     uint(rand.Intn(1000)),    // Random ItemID
		ItemNumber: uuid.New(),               // Random UUID
		Name:       generateRandomString(10), // Random 10-character string
		Selected:   rand.Intn(2) == 0,        // Randomly true or false
	}
}

// Helper function to generate a random string of a given length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateItemFixtures generates multiple Item fixtures
func GenerateItemFixtures(count int) []Item {
	rand.Seed(uint64(time.Now().UnixNano()))
	items := make([]Item, count)
	for i := 0; i < count; i++ {
		items[i] = GenerateItemFixture()
	}
	return items
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

// MockGormDB is a mock for ItemDAO to mimic transaction related to Item in the database
type MockItemDAO struct {
	mock.Mock
}

func (m *MockItemDAO) GetAllItems() ([]Item, error) {
	args := m.Called()
	return args.Get(0).([]Item), args.Error(1)
}
