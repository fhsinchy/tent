package services

import (
	"github.com/fhsinchy/tent/types"
)

// Neo4j service holds necessary data for creating and running the Neo4j container.
var Neo4j types.Service = types.Service{
	Name:  "neo4j",
	Image: "docker.io/neo4j",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "neo4j-data",
			Dest: "/data",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "HTTP Port",
			ContainerPort: 7474,
			HostPort:      7474,
		},
		{
			Text:          "Bolt Port",
			ContainerPort: 7687,
			HostPort:      7687,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Authentication (user/pass)",
			Key:     "NEO4J_AUTH",
			Value:   "neo4j/secret",
			Mutable: true,
		},
	},
	InsecureEnv: []types.EnvVar{
		{Key: "NEO4J_AUTH", Value: "none"},
	},
	InsecureInfo: "authentication disabled, no credentials required",
}
