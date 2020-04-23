package event

import "context"

type (
	// Type uniquely identifies an event's type
	// e.g. world.block.placed
	Type string

	// Event represents the basic data of an event
	Event interface {
		Type() Type
		Context() context.Context
	}

	// Router handles the routing of events
	Router interface {
		Push(Event)
		On(Type, func(Event))
	}
)
