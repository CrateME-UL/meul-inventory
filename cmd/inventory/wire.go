//go:build wireinject
// +build wireinject

package main

import (
	"meul/inventory/internal/infrastructures"
	event "meul/inventory/internal/infrastructures/models"

	"github.com/google/wire"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeEvent(config infrastructures.DbConfig) (event.Event, error) {
	wire.Build(event.NewEvent, infrastructures.NewDatabaseConnection, event.NewGreeter, event.NewMessage)
	return event.Event{}, nil
}
