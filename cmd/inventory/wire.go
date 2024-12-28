//go:build wireinject
// +build wireinject

package main

import (
	infrastructures_drivers_postgres "meul/inventory/internal/infrastructures/drivers/postgres"
	event "meul/inventory/internal/infrastructures/drivers/postgres/models"

	"github.com/google/wire"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeEvent(config infrastructures_drivers_postgres.DbConfig) (event.Event, error) {
	wire.Build(event.NewEvent, infrastructures_drivers_postgres.NewDatabaseConnection, event.NewGreeter, event.NewMessage)
	return event.Event{}, nil
}
