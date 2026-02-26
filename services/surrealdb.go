package services

import (
	"github.com/fhsinchy/tent/types"
)

// SurrealDB service holds necessary data for creating and running the SurrealDB container.
var SurrealDB types.Service = types.Service{
	Name:  "surrealdb",
	Image: "docker.io/surrealdb/surrealdb",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "surrealdb-data",
			Dest: "/data",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 8000,
			HostPort:      8000,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Root Username",
			Key:     "SURREAL_USER",
			Value:   "root",
			Mutable: true,
		},
		{
			Text:    "Root Password",
			Key:     "SURREAL_PASS",
			Value:   "secret",
			Mutable: true,
		},
	},
	InsecureEnv:  []types.EnvVar{},
	InsecureInfo: "authentication disabled, no credentials required",
	Command:      []string{"start", "--log", "info", "file:/data/srdb.db"},
}
