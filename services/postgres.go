package services

import (
	"github.com/fhsinchy/tent/types"
)

// Postgres service holds necessary data for creating and running the Postgres container.
var Postgres types.Service = types.Service{
	Name:  "postgres",
	Image: "docker.io/postgres",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "postgres-data",
			Dest: "/var/lib/postgresql/data",
		},
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
}
