package docker

import (
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

type Image struct {
	Name string
}

type InitImage interface {
	EnsureImage(string) error
}

func (c *Controller) EnsureImage(image string) (err error) {
	reader, err := c.cli.ImagePull(c.ctx, image, types.ImagePullOptions{})

	if err != nil {
		return err
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
	return nil
}
