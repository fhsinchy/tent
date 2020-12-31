package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// Mongo service holds necessary data for creating and running the Mongo container.
var Mongo types.Service = types.Service{
	Name:  "mongo",
	Image: "docker.io/mongo",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/data/db",
	},
	PortMappings: []types.PortMapping{
		{
			Text: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 27017,
				HostPort:      27017,
			},
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Server Admin Username",
			Key:     "MONGO_INITDB_ROOT_USERNAME",
			Value:   "admin",
			Mutable: true,
		},
		{
			Text:    "Server Admin Password",
			Key:     "MONGO_INITDB_ROOT_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
	},
	Command:    []string{"--serviceExecutor", "adaptive"},
	HasVolumes: true,
	HasCommand: true,
}
