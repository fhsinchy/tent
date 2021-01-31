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
			Text:          "Server Port",
			ContainerPort: 5432,
			HostPort:      5432,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Server Root Password",
			Key:     "POSTGRES_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
	},
	HasVolumes: true,
}
