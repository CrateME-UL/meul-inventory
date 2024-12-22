//go:build wireinject
// +build wireinject

package main

import (
	event "meul/inventory/internal/infrastructures"

	"github.com/google/wire"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeEvent(phrase string) (event.Event, error) {
	wire.Build(event.NewEvent, event.NewGreeter, event.NewMessage)
	return event.Event{}, nil
}
