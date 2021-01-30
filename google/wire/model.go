package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/wire"
)

// Message Message
type Message string

// NewMessage NewMessage
func NewMessage(m string) Message {
	return Message(m)
}

// NewGreeter NewGreeter
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// Greeter Greeter
type Greeter struct {
	Message Message
	Grumpy  bool
}

// Greet Greet
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// NewEvent NewEvent
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

//Event Event
type Event struct {
	Greeter Greeter
}

// Start Start
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// MegaSet MegaSet
var MegaSet = wire.NewSet(NewEvent, NewGreeter, NewMessage)
