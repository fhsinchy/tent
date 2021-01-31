package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// MeiliSearch service holds necessary data for creating and running the MeiliSearch container.
var MeiliSearch types.Service = types.Service{
	Name:  "meilisearch",
	Image: "docker.io/getmeili/meilisearch",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/data.ms",
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 7700,
			HostPort:      7700,
		},
	},
	HasVolumes: true,
}
