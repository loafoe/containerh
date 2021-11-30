package containerh

import (
	"net/http"

	"github.com/philips-software/go-hsdp-api/cartel"
)

func NewClient(httpClient *http.Client, config Config) (*Client, error) {
	c := &Client{
		config: config,
	}
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
	c.cartelClient = cartelClient
	c.Instances = &InstancesService{Client: c}
	return c, nil
}
