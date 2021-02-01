package services

import (
	"github.com/fhsinchy/tent/types"
)

// PostGIS service holds necessary data for creating and running the PostGIS container.
var PostGIS types.Service = types.Service{
	Name:  "postgis",
	Image: "docker.io/postgis/postgis",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "postgis-data",
			Dest: "/var/lib/postgis/data",
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
