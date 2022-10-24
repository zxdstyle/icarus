package buffer

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type (
	Buffer[V any] struct {
		buffer    []chan V
		shardFunc shardFunc[V]
		do        doFunc[V]
		wait      sync.WaitGroup
		opts      options
	}

	doFunc[V any] func(ctx context.Context, val []V)

	shardFunc[V any] func(val V) int
)

func New[V any](do doFunc[V], shard shardFunc[V], opts ...Option) *Buffer[V] {
	b := &Buffer[V]{
		do:        do,
		shardFunc: shard,
	}
	for _, opt := range opts {
		opt.apply(&b.opts)
	}
	b.opts.check()

	b.buffer = make([]chan V, b.opts.worker)
	for i := 0; i < b.opts.worker; i++ {
		b.buffer[i] = make(chan V, b.opts.buffer)
	}

	go b.start()

	return b
}

func (b *Buffer[V]) start() {
	b.wait.Add(len(b.buffer))
	for i := range b.buffer {
		go b.merge(b.buffer[i])
	}
}

func (b *Buffer[V]) Push(values ...V) {
	for i := range values {
		val := values[i]
		sharding := b.shardFunc(val)
		ch := b.buffer[sharding]
		ch <- val
	}
}

func (b *Buffer[V]) merge(ch chan V) {
	defer b.wait.Done()

	var (
		count int
		quit  bool
		vals  = make([]V, 0)
	)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ticker := time.NewTicker(b.opts.interval)
	for {
		select {
		case <-ticker.C:
			if len(vals) > 0 {
				b.doFunc(&vals, &count)
			}

			if quit && len(vals) == 0 && len(ch) == 0 {
				return
			}
		case msg := <-ch:
			count++
			vals = append(vals, msg)
			if count >= b.opts.limit {
				b.doFunc(&vals, &count)
			}

			if quit && len(vals) == 0 && len(ch) == 0 {
				return
			}
		case <-sc:
			quit = true
			if len(vals) == 0 && len(ch) == 0 {
				return
			}
		}
	}
}

func (b *Buffer[V]) doFunc(vals *[]V, count *int) {
	ctx := context.Background()
	b.do(ctx, *vals)
	*vals = make([]V, 0)
	*count = 0
}

func (b *Buffer[V]) WaitDone() {
	b.wait.Wait()
}
