package containerh

import (
	"fmt"
	"log"
	"net/http"
	"os"

	docker "github.com/docker/docker/client"
)

func NewClient(httpClient *http.Client, config Config) (*Client, error) {
	c := &Client{
		config: config,
	}
	if config.DebugLog != "" {
		f, err := os.OpenFile(config.DebugLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("error opening file: %v", err)
		} else {
			log.SetOutput(f)
		}
	}
	/*
		cartelClient, err := cartel.NewClient(httpClient, &cartel.Config{
			Region:     config.Region,
			Token:      config.CartelKey,
			Secret:     config.CartelSecret,
			DebugLog:   config.DebugLog,
			SkipVerify: true,
		})
		if err != nil {
			return nil, err
		}
	*/
	// Docker client
	dockerClient, err := docker.NewClientWithOpts(docker.FromEnv)
	if err != nil {
		fmt.Printf("Error creating client (%s): %v\n", os.Getenv("DOCKER_HOST"), err)
		return nil, err
	}

	//c.cartelClient = cartelClient
	c.Instances = &InstancesService{Client: c}
	c.CLI = &CLIServices{Client: c, LocalDockerClient: dockerClient}
	return c, nil
}
