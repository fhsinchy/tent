package services

import (
	"github.com/fhsinchy/tent/types"
)

// Typesense service holds necessary data for creating and running the Typesense container.
var Typesense types.Service = types.Service{
	Name:  "typesense",
	Image: "docker.io/typesense/typesense",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "typesense-data",
			Dest: "/data",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 8108,
			HostPort:      8108,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "API Key",
			Key:     "TYPESENSE_API_KEY",
			Value:   "secret",
			Mutable: true,
		},
		{
			Text:    "Data Directory",
			Key:     "TYPESENSE_DATA_DIR",
			Value:   "/data",
			Mutable: false,
		},
	},
}
