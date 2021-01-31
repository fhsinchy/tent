package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// ElasticSearch service holds necessary data for creating and running the ElasticSearch container.
var ElasticSearch types.Service = types.Service{
	Name:  "elasticsearch",
	Image: "docker.io/elasticsearch",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/usr/share/elasticsearch/data",
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 9200,
			HostPort:      9200,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Discovery Mode",
			Key:     "discovery.type",
			Value:   "single-node",
			Mutable: false,
		},
	},
	HasVolumes: true,
}
