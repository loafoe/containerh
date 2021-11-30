package docker

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/philips-software/go-hsdp-api/cartel"
)

type ImageConfig struct {
	Image    string `json:"dockerImage"`
	Username string `json:"dockerUsername"`
	Password string `json:"dockerPassword"`
}

type Container struct {
	ClusterID         string
	LocalDockerClient *client.Client
	EnableFluentd     bool
	Cartel            *cartel.Client
}

func (c *Container) GetAllInstances(ctx context.Context) ([]cartel.InstanceDetails, error) {
	instances, _, err := c.Cartel.GetAllInstances()
	if instances == nil {
		return []cartel.InstanceDetails{}, err
	}
	return *instances, err
}

func (c *Container) GetWorkerInfo(ctx context.Context) (types.Info, error) {
	info, err := c.LocalDockerClient.Info(ctx)
	if err != nil {
		return types.Info{}, fmt.Errorf("dockerClient.Info: %w", err)
	}
	return info, nil
}

func (c *Container) ContainerList(ctx context.Context) ([]types.Container, error) {
	return c.LocalDockerClient.ContainerList(context.Background(), types.ContainerListOptions{})
}

func (c *Container) ContainerKill(ctx context.Context, id string) error {
	return c.LocalDockerClient.ContainerKill(context.Background(), id, "")
}

func (c *Container) DockerLogin(ctx context.Context, config ImageConfig) ([]byte, error) {
	ref, err := reference.ParseNormalizedNamed(config.Image)
	if err != nil {
		return []byte{}, err
	}
	registry := ""
	if str := strings.Split(ref.Name(), "/"); len(str) > 1 {
		registry = str[0]
	}
	cmd := exec.Command("docker", "login", registry, "--username", config.Username, "--password", config.Password)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}
	return output, err
}

func (c *Container) DockerCommand(ctx context.Context, args []string) ([]byte, error) {
	path, err := exec.LookPath("docker")
	if err != nil {
		return []byte{}, err
	}
	cmd := exec.Command(path, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return []byte{}, err
	}
	return output, err
}
