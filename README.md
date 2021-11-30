# containerh

A client to wrap / hide the subtleties of the HSP Container Host environment

## Usage

```golang
package main

import (
	"fmt"
	"net/http"

	"github.com/loafoe/containerh"
	"github.com/philips-software/go-hsdp-api/cartel"
)

func main() {
	client, err := containerh.NewClient(http.DefaultClient, containerh.Config{
		Region:       "eu-west",
		CartelKey:    "key",
		CartelSecret: "secret",
		Username:     "jholden",
		SSHAgent:     true,
	})

	instances, _, err := client.Instances.All()

	if err != nil {
		fmt.Printf("Error retrieving instances: %v\n", err)
		return
	}

	for _, inst := range *instances {
		fmt.Printf("instance: %s\n", inst.NameTag)
	}
}
```

# license

License is MIT
