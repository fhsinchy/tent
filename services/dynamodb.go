package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// DynamoDB service holds necessary data for creating and running the DynamoDB container.
var DynamoDB types.Service = types.Service{
	Name:  "dynamodb",
	Image: "docker.io/amazon/dynamodb-local",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/dynamodb_local_db",
	},
	PortMappings: []types.PortMapping{
		{
			Text: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 8000,
				HostPort:      8000,
			},
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
	Command:    []string{"-jar", "DynamoDBLocal.jar", "--sharedDb", "-dbPath", "/dynamodb_local_db"},
	HasVolumes: true,
	HasCommand: true,
}
