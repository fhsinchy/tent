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
	PortMappings: []specgen.PortMapping{
		{
			ContainerPort: 8081,
			HostPort:      8081,
		},
	},
	Env: map[string]string{
		"ME_CONFIG_OPTIONS_EDITORTHEME": "ambiance",
	},
	Prompts: map[string]bool{
		"tag":  true,
		"port": true,
	},
}
