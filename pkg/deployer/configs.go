package deployer

import (
	"github.com/rajaSahil/deployer/pkg/deployer/dependencies"
	"github.com/rajaSahil/deployer/pkg/deployer/docker"
)

type Config struct {
	Controller   docker.Controller
	Dependencies dependencies.Dependencies
}
