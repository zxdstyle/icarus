package events

import (
	"context"
	"fmt"
)

var defaultRetryTimes = 3

func Subscribe[T any](bus *Bus, event Event[T], subscribers ...Subscriber[T]) {
	name := resolveEventName(event)
	for idx, _ := range subscribers {
		bus.set(name, subscribers[idx])
	}
}

func Dispatch[T any](ctx context.Context, bus *Bus, event Event[T]) error {
	name := resolveEventName(event)
	if !bus.exists(name) {
		return nil
	}
	return doDispatchEvent(ctx, bus, name, event)
}

func Shutdown(bus *Bus) {
	// todo waiting log
	bus.wait()
}

func doDispatchEvent[T any](ctx context.Context, bus *Bus, name string, e Event[T]) error {
	subscribers, found := bus.get(name)
	if !found || subscribers == nil || subscribers.IsEmpty() {
		return nil
	}

	subscribers.Iterator(func(subscriber any) bool {
		sub, ok := subscriber.(Subscriber[T])
		if !ok {
			return true
		}
		handleDispatch(ctx, bus, sub, e)
		return true
	})
	return nil
}

func handleDispatch[T any](ctx context.Context, bus *Bus, subscriber Subscriber[T], e Event[T]) {
	if asyncSub, is := subscriber.(AsyncSubscriber); is && asyncSub.Async() {
		bus.add(1)
		go func() {
			defer bus.done()
			doHandle(context.Background(), subscriber, e)
		}()
		return
	}
	doHandle(ctx, subscriber, e)
}

func doHandle[T any](ctx context.Context, subscriber Subscriber[T], e Event[T]) {
	retry := defaultRetryTimes
	val, ok := subscriber.(SpecifiedRetryTimes)
	if ok {
		retry = val.Retry()
	}

	if retry == 0 {
		err := subscriber.Handle(ctx, e.Payload())
		if err != nil {
			if handler, ok := subscriber.(FailedSubscriber[T]); ok {
				handler.OnFailed(ctx, e, err)
			}
		}
		return
	}

	times := 0
	for {
		err := subscriber.Handle(ctx, e.Payload())
		if err == nil {
			break
		}

		if times >= retry {
			if handler, ok := subscriber.(FailedSubscriber[T]); ok {
				handler.OnFailed(ctx, e, err)
			}
			break
		}
		times++
	}
}

func resolveEventName[T any](e Event[T]) string {
	event, ok := e.(NamedEvent)
	if ok {
		name := event.Name()
		if len(name) > 0 {
			return name
		}
	}
	return fmt.Sprintf("%T", e)
}
