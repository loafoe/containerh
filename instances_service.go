package containerh

import (
	"fmt"

	"github.com/philips-software/go-hsdp-api/cartel"
)

type InstancesService struct {
	*Client
}

func (i *InstancesService) All() (*[]cartel.InstanceDetails, *cartel.Response, error) {
	if i.cartelClient == nil {
		return nil, nil, fmt.Errorf("cartelClient not initalised")
	}
	return i.cartelClient.GetAllInstances()
}
