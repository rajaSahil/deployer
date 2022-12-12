package dependencies

import (
	"github.com/docker/docker/api/types/container"
	"github.com/rajaSahil/deployer/pkg/deployer/docker"
)

type Pulsar struct {
	*DependencyConfigs
}

func (p *Pulsar) InitializeDependencyConfigs(enabled bool, addr string) {
	p.Enable = enabled
	p.Address = addr
}

func (p *Pulsar) InitializeImage(img *docker.Image) {
	p.image = *img
}

func (p *Pulsar) InitializeContainerConfig(ctrCfg *docker.Container) {
	p.ctr = *ctrCfg
}

func (p *Pulsar) GetImage() *docker.Image {
	return &p.image
}

func (p *Pulsar) GetContainerConfig() *docker.Container {
	return &p.ctr
}

func (d *Dependencies) InitPulsar() *Pulsar {
	d.Pulsar.InitializeImage(&docker.Image{
		Name: "pulsar",
	}) // Call it from configs
	d.Pulsar.InitializeContainerConfig(&docker.Container{
		Cfg: &container.Config{
			Hostname:        "",
			Domainname:      "",
			User:            "",
			AttachStdin:     false,
			AttachStdout:    false,
			AttachStderr:    false,
			ExposedPorts:    nil,
			Tty:             false,
			OpenStdin:       false,
			StdinOnce:       false,
			Env:             nil,
			Cmd:             nil,
			Healthcheck:     nil,
			ArgsEscaped:     false,
			Image:           "",
			Volumes:         nil,
			WorkingDir:      "",
			Entrypoint:      nil,
			NetworkDisabled: false,
			MacAddress:      "",
			OnBuild:         nil,
			Labels:          nil,
			StopSignal:      "",
			StopTimeout:     nil,
			Shell:           nil,
		},
	})
	return &d.Pulsar
}
