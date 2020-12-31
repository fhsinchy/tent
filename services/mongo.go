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
			Name: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 27017,
				HostPort:      27017,
			},
		},
	},
	Env: []types.EnvVar{
		{
			Name:    "Server Admin Username",
			Key:     "MONGO_INITDB_ROOT_USERNAME",
			Value:   "admin",
			Mutable: true,
		},
		{
			Name:    "Server Admin Password",
			Key:     "MONGO_INITDB_ROOT_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
	},
	Command:    []string{"--serviceExecutor", "adaptive"},
	HasVolumes: true,
	HasCommand: true,
	Prompts: map[string]bool{
		"tag":      true,
		"password": true,
		"volume":   true,
		"port":     true,
		"username": true,
	},
}
