package services

import (
	"github.com/fhsinchy/tent/types"
)

// MeiliSearch service holds necessary data for creating and running the MeiliSearch container.
var MeiliSearch types.Service = types.Service{
	Name:  "meilisearch",
	Image: "docker.io/getmeili/meilisearch",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "meilisearch-data",
			Dest: "/data.ms",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 7700,
			HostPort:      7700,
		},
	},
}
