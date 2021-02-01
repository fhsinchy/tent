package services

import (
	"github.com/fhsinchy/tent/types"
)

// Redis service holds necessary data for creating and running the Redis container.
var Redis types.Service = types.Service{
	Name:  "redis",
	Image: "docker.io/redis",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "redis-data",
			Dest: "/data",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 6379,
			HostPort:      6379,
		},
	},
}
