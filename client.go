package containerh

import (
	"github.com/philips-software/go-hsdp-api/cartel"
)

type Client struct {
	cartelClient *cartel.Client
	config       Config

	Instances *InstancesService
}
