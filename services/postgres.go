package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// Postgres service holds necessary data for creating and running the Postgres container.
var Postgres types.Service = types.Service{
	Name:  "postgres",
	Image: "docker.io/postgres",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/var/lib/postgresql/data",
	},
	PortMappings: []types.PortMapping{
		{
			Name: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 5432,
				HostPort:      5432,
			},
		},
	},
	Env: []types.EnvVar{
		{
			Name:    "Server Root Password",
			Key:     "POSTGRES_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
	},
	HasVolumes: true,
	Prompts: map[string]bool{
		"tag":      true,
		"password": true,
		"volume":   true,
		"port":     true,
	},
}
