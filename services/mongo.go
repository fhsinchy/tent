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
	PortMapping: specgen.PortMapping{
		ContainerPort: 27017,
		HostPort:      27017,
	},
	Env: map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": "admin",
		"MONGO_INITDB_ROOT_PASSWORD": "secret",
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
