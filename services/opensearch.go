package services

import (
	"github.com/fhsinchy/tent/types"
)

// OpenSearch service holds necessary data for creating and running the OpenSearch container.
var OpenSearch types.Service = types.Service{
	Name:  "opensearch",
	Image: "docker.io/opensearchproject/opensearch",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "opensearch-data",
			Dest: "/usr/share/opensearch/data",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "REST API Port",
			ContainerPort: 9200,
			HostPort:      9200,
		},
		{
			Text:          "Performance Analyzer Port",
			ContainerPort: 9600,
			HostPort:      9600,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Discovery Mode",
			Key:     "discovery.type",
			Value:   "single-node",
			Mutable: false,
		},
		{
			Text:    "Disable Security Plugin",
			Key:     "DISABLE_SECURITY_PLUGIN",
			Value:   "true",
			Mutable: false,
		},
	},
}
