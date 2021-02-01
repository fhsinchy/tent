package services

import (
	"github.com/fhsinchy/tent/types"
)

// Mongo service holds necessary data for creating and running the Mongo container.
var Mongo types.Service = types.Service{
	Name:  "mongo",
	Image: "docker.io/mongo",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "mongo-data",
			Dest: "/data/db",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 27017,
			HostPort:      27017,
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
	Command: []string{"--serviceExecutor", "adaptive"},
}
