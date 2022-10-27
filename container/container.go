package container

import "github.com/samber/do"

func New() *do.Injector {
	return do.New()
}

type Provider[T any] func() (T, error)

func Provide[T any](injector *do.Injector, provider Provider[T]) {
	do.Provide(injector, wrapProvider(provider))
}

func ProvideValue[T any](injector *do.Injector, value T) {
	do.ProvideValue(injector, value)
}

func ProvideNamedValue[T any](injector *do.Injector, name string, value T) {
	do.ProvideNamedValue(injector, name, value)
}

func ProvideNamed[T any](injector *do.Injector, name string, provider Provider[T]) {
	do.ProvideNamed(injector, name, wrapProvider(provider))
}

func Invoke[T any](injector *do.Injector) (T, error) {
	return do.Invoke[T](injector)
}

func InvokeNamed[T any](injector *do.Injector, name string) (T, error) {
	return do.InvokeNamed[T](injector, name)
}

func MustInvoke[T any](injector *do.Injector) T {
	return do.MustInvoke[T](injector)
}

func MustInvokeNamed[T any](injector *do.Injector, name string) T {
	return do.MustInvokeNamed[T](injector, name)
}

func wrapProvider[T any](provider Provider[T]) do.Provider[T] {
	return func(i *do.Injector) (T, error) {
		return provider()
	}
}
