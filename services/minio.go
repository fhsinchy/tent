package services

import (
	"github.com/fhsinchy/tent/types"
)

// MinIO service holds necessary data for creating and running the MinIO container.
var MinIO types.Service = types.Service{
	Name:  "minio",
	Image: "docker.io/minio/minio",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "minio-data",
			Dest: "/data",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "API Port",
			ContainerPort: 9000,
			HostPort:      9000,
		},
		{
			Text:          "Console Port",
			ContainerPort: 9001,
			HostPort:      9001,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Root Username",
			Key:     "MINIO_ROOT_USER",
			Value:   "minioadmin",
			Mutable: true,
		},
		{
			Text:    "Root Password",
			Key:     "MINIO_ROOT_PASSWORD",
			Value:   "minioadmin",
			Mutable: true,
		},
	},
	Command: []string{"server", "/data", "--console-address", ":9001"},
}
