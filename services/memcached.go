package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// Memcached service holds necessary data for creating and running the Memcached container.
var Memcached types.Service = types.Service{
	Name:  "memcached",
	Image: "docker.io/memcached",
	Tag:   "latest",
	PortMapping: specgen.PortMapping{
		ContainerPort: 11211,
		HostPort:      11211,
	},
	Env: map[string]string{},
	Prompts: map[string]bool{
		"tag":  true,
		"port": true,
	},
}
