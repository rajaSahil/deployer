package dependencies

import (
	"github.com/rajaSahil/deployer/pkg/deployer/docker"
)

type Dependencies struct {
	Enabled bool
	Pulsar  Pulsar
	Mongo   Mongo
	Mysql   Mysql
}

type DependencyConfigs struct {
	Enable  bool
	Address string
	image   docker.Image
	ctr     docker.Container
}

type InitDependency interface {
	InitializeDependencyConfigs(bool, string)
	InitializeImage(*docker.Image)
	InitializeContainerConfig(*docker.Container)
	GetImage() *docker.Image
	GetContainerConfig() *docker.Container
}
