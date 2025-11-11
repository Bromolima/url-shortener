package injector

import (
	"go.uber.org/dig"
)

func Provide(container *dig.Container, constructor any, opts ...dig.ProvideOption) error {
	if err := container.Provide(constructor, opts...); err != nil {
		return err
	}

	return nil
}

func Resolve[T any](container *dig.Container) (T, error) {
	var invoked T
	if err := container.Invoke(func(invk T) {
		invoked = invk
	}); err != nil {
		return invoked, err
	}

	return invoked, nil
}
