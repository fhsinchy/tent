package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// Redis service holds necessary data for creating and running the Redis container.
var Redis types.Service = types.Service{
	Name:  "redis",
	Image: "docker.io/redis",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/data",
	},
	PortMappings: []types.PortMapping{
		{
			Text: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 6379,
				HostPort:      6379,
			},
		},
	},
	HasVolumes: true,
	Prompts: map[string]bool{
		"tag":    true,
		"volume": true,
		"port":   true,
	},
}
