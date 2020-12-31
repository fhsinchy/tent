package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// PostGIS service holds necessary data for creating and running the PostGIS container.
var PostGIS types.Service = types.Service{
	Name:  "postgis",
	Image: "docker.io/postgis/postgis",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/var/lib/postgis/data",
	},
	PortMappings: []types.PortMapping{
		{
			Text: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 5432,
				HostPort:      5432,
			},
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
	Prompts: map[string]bool{
		"tag":      true,
		"password": true,
		"volume":   true,
		"port":     true,
	},
}
