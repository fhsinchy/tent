package services

import (
	"github.com/fhsinchy/tent/types"
)

// DynamoDB service holds necessary data for creating and running the DynamoDB container.
var DynamoDB types.Service = types.Service{
	Name:  "dynamodb",
	Image: "docker.io/amazon/dynamodb-local",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "dynamodb-data",
			Dest: "/dynamodb_local_db",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 8000,
			HostPort:      8000,
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
	Command: []string{"-jar", "DynamoDBLocal.jar", "--sharedDb", "-dbPath", "/dynamodb_local_db"},
}
