package models

import (
	"fmt"

	"gorm.io/gorm"
)

type ItemID string

type Greeter struct {
	Message ItemID
	DB      *gorm.DB
}

type Event struct {
	Greeter Greeter
}

func NewMessage() ItemID {
	return ItemID("Hi there!")
}

func NewGreeter(m ItemID, db *gorm.DB) Greeter {
	return Greeter{Message: m, DB: db}
}

func (g Greeter) Greet() ItemID {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
