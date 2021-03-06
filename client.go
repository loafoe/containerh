package containerh

import (
	"context"
	"fmt"
	"log"

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
	log.Printf("Pull image: %s\n", imageName)
	log.Printf("opts: %v\n", opts)
	return nil, fmt.Errorf("pull not implemented yet")
}

func (c *Client) NewContainer(ctx context.Context, containerName string, runtime containerd.NewContainerOpts, snapshot containerd.NewContainerOpts, spec containerd.NewContainerOpts) (containerd.Container, error) {
	log.Printf("New container: %s\n", containerName)
	return nil, fmt.Errorf("newContainer not implemented yet")
}

func (c *Client) LoadContainer(ctx context.Context, id string) (containerd.Container, error) {
	log.Printf("Load container: %s\n", id)
	return nil, fmt.Errorf("loadContainer not implemented yet")
}
