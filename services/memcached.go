package services

import (
	"github.com/fhsinchy/tent/types"
)

// Memcached service holds necessary data for creating and running the Memcached container.
var Memcached types.Service = types.Service{
	Name:  "memcached",
	Image: "docker.io/memcached",
	Tag:   "latest",
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 11211,
			HostPort:      11211,
		},
	},
}
