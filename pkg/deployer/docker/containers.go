package docker

import (
	"github.com/docker/docker/api/types/container"
)

type Container struct {
	Cfg *container.Config
}

type InitContainer interface {
	EnsureContainer(string, *container.Config) *Container
}
