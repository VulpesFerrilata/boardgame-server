package container

import (
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	return container
}
