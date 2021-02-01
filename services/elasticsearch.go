package services

import (
	"github.com/fhsinchy/tent/types"
)

// ElasticSearch service holds necessary data for creating and running the ElasticSearch container.
var ElasticSearch types.Service = types.Service{
	Name:  "elasticsearch",
	Image: "docker.io/elasticsearch",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "elasticsearch-data",
			Dest: "/usr/share/elasticsearch/data",
		},
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
}
