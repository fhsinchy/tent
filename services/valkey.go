package services

import (
	"github.com/fhsinchy/tent/types"
)

// Valkey service holds necessary data for creating and running the Valkey container.
var Valkey types.Service = types.Service{
	Name:  "valkey",
	Image: "docker.io/valkey/valkey",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "valkey-data",
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
