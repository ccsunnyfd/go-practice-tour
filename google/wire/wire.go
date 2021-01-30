//+build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeEvent(m string) (Event, error) {
	// wire.Build(NewEvent, NewGreeter, NewMessage)
	wire.Build(MegaSet)
	return Event{}, nil
}
