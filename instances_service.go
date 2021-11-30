package containerh

import "github.com/philips-software/go-hsdp-api/cartel"

type InstancesService struct {
	*Client
}

func (i *InstancesService) All() (*[]cartel.InstanceDetails, *cartel.Response, error) {
	return i.cartelClient.GetAllInstances()
}
