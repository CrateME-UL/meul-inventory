package event

import (
	"fmt"

	"gorm.io/gorm"
)

type Message string

type Greeter struct {
	Message Message
	DB      *gorm.DB
}

type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message, db *gorm.DB) Greeter {
	return Greeter{Message: m, DB: db}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
