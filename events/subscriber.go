package events

import "context"

type (
	SubscriberSorted struct {
		subscribers []any
	}

	Subscriber[T any] interface {
		Handle(ctx context.Context, payload T) error
	}

	AsyncSubscriber interface {
		Async() bool
	}

	FailedSubscriber[V any] interface {
		OnFailed(ctx context.Context, e Event[V], err error)
	}

	SpecifiedRetryTimes interface {
		Retry() int
	}
)

func (l *SubscriberSorted) IsEmpty() bool {
	return l == nil || l.Size() == 0
}

func (l *SubscriberSorted) Size() int {
	return len(l.subscribers)
}

func (l *SubscriberSorted) Push(subscribers ...any) *SubscriberSorted {
	l.subscribers = append(l.subscribers, subscribers...)
	return l
}

func (l *SubscriberSorted) Iterator(fn func(subscriber any) bool) {
	for _, subscriber := range l.subscribers {
		if !fn(subscriber) {
			break
		}
	}
}
