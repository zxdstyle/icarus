package events

type (
	Event[T any] interface {
		Payload() T
	}

	NamedEvent interface {
		Name() string
	}
)
