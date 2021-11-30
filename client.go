package containerh

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/philips-software/go-hsdp-api/cartel"
)

type Client struct {
	cartelClient *cartel.Client
	config       Config

	Instances *InstancesService
	CLI       *CLIServices
}

// Version of containerd
type Version struct {
	// Version number
	Version string
	// Revision from git that was built
	Revision string
}

func (c *Client) IsServing(ctx context.Context) (bool, error) {
	// TODO: implement properly
	return true, nil
}

func (c *Client) Version(ctx context.Context) (Version, error) {
	return Version{
		Version:  "0.0.1",
		Revision: "deadbeef",
	}, nil
}

func (c *Client) Pull(ctx context.Context, imageName string, opts ...containerd.RemoteOpt) (containerd.Image, error) {
	return nil, nil
}

func (c *Client) NewContainer(ctx context.Context, containerName string, runtime containerd.NewContainerOpts, snapshot containerd.NewContainerOpts, spec containerd.NewContainerOpts) (containerd.Container, error) {
	return nil, nil
}

func (c *Client) LoadContainer(ctx context.Context, id string) (containerd.Container, error) {
	return nil, nil
}
