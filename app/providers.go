package app

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/zxdstyle/icarus/container"
	"github.com/zxdstyle/icarus/database"
	"github.com/zxdstyle/icarus/events"
	"github.com/zxdstyle/icarus/server"
	"github.com/zxdstyle/icarus/server/engines"
	"github.com/zxdstyle/icarus/server/engines/fiber"
	"gorm.io/gorm"
)

var (
	providers       = container.New()
	defaultEventBus = events.NewBus()
)

func init() {
	config.AddDriver(yamlv3.Driver)
	if err := config.LoadFiles("config.yaml"); err != nil {
		panic(err)
	}

	container.Provide(providers, fiber.NewFiber)

	container.Provide(providers, func() (*server.Server, error) {
		engine := container.MustInvoke[engines.Engine](providers)
		return server.New(engine), nil
	})

	container.Provide(providers, func() (*gorm.DB, error) {
		var cfg database.Config
		if err := config.BindStruct("database", &cfg); err != nil {
			return nil, err
		}
		return database.Connect(cfg)
	})

	container.Provide(providers, func() (*redis.Client, error) {
		var cfg redis.Options
		if err := config.BindStruct("redis", &cfg); err != nil {
			return nil, err
		}
		return database.NewRedis(&cfg), nil
	})
}

func Make[T any](name ...string) T {
	if len(name) > 0 {
		return container.MustInvokeNamed[T](providers, name[0])
	}
	return container.MustInvoke[T](providers)
}

func Server() *server.Server {
	return container.MustInvoke[*server.Server](providers)
}

func DB() *gorm.DB {
	return container.MustInvoke[*gorm.DB](providers)
}

func Redis() *redis.Client {
	return container.MustInvoke[*redis.Client](providers)
}

func Subscriber[T any](event events.Event[T], subscribers ...events.Subscriber[T]) {
	events.Subscribe[T](defaultEventBus, event, subscribers...)
}

func Dispatch[T any](ctx context.Context, event events.Event[T]) error {
	return events.Dispatch(ctx, defaultEventBus, event)
}

func Shutdown() {
	events.Shutdown(defaultEventBus)
}
