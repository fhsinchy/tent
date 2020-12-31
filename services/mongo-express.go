package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// MongoExpress service holds necessary data for creating and running the MongoExpress container.
var MongoExpress types.Service = types.Service{
	Name:  "mongo-express",
	Image: "docker.io/mongo-express",
	Tag:   "latest",
	PortMappings: []types.PortMapping{
		{
			Text: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 8081,
				HostPort:      8081,
			},
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Editor UI Theme",
			Key:     "ME_CONFIG_OPTIONS_EDITORTHEME",
			Value:   "ambiance",
			Mutable: false,
		},
	},
	Prompts: map[string]bool{
		"tag":  true,
		"port": true,
	},
}
