package events

import (
	"github.com/zxdstyle/liey/pkg/co"
	"sync"
)

type Bus struct {
	subscribers     *co.Map[string, *SubscriberSorted]
	subscribedNames *co.Map[string, int]
	waitGroup       sync.WaitGroup
}

func NewBus() *Bus {
	return &Bus{
		subscribers:     co.NewMap[string, *SubscriberSorted](),
		subscribedNames: co.NewMap[string, int](),
		waitGroup:       sync.WaitGroup{},
	}
}

func (b *Bus) exists(name string) bool {
	_, ok := b.subscribers.Search(name)
	return ok
}

func (b *Bus) get(name string) (sub *SubscriberSorted, found bool) {
	sub, found = b.subscribers.Search(name)
	return
}

func (b *Bus) set(name string, subscribers ...any) {
	item, found := b.subscribers.Search(name)
	if !found || item.IsEmpty() {
		b.subscribers.Set(name, &SubscriberSorted{subscribers: subscribers})
		return
	}
	item.Push(subscribers...)
}

func (b *Bus) add(delta int) {
	b.waitGroup.Add(delta)
}

func (b *Bus) done() {
	b.waitGroup.Done()
}

func (b *Bus) wait() {
	b.waitGroup.Wait()
}
