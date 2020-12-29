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
	PortMapping: specgen.PortMapping{
		ContainerPort: 3306,
		HostPort:      3306,
	},
	Env: map[string]string{
		"POSTGRES_PASSWORD": "secret",
	},
	HasVolumes: true,
	Prompts: map[string]bool{
		"tag":      true,
		"password": true,
		"volume":   true,
		"port":     true,
	},
}
