package docker

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/rajaSahil/deployer/pkg/deployer/dependencies"
)

type Controller struct {
	ctx context.Context
	cli *client.Client
	dep dependencies.Dependencies
}

func NewController() (*Controller, error) {
	c := Controller{
		ctx: context.Background(),
		cli: nil,
	}
	var err error
	c.cli, err = client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func Run() {
	//_, err := NewController()
	//
	//if err != nil {
	//	//klog.Info("alkndcl")
	//	return
	//}

}
