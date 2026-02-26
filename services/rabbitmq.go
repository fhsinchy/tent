package services

import (
	"github.com/fhsinchy/tent/types"
)

// RabbitMQ service holds necessary data for creating and running the RabbitMQ container.
var RabbitMQ types.Service = types.Service{
	Name:  "rabbitmq",
	Image: "docker.io/rabbitmq",
	Tag:   "management",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "rabbitmq-data",
			Dest: "/var/lib/rabbitmq",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 5672,
			HostPort:      5672,
		},
		{
			Text:          "Management UI Port",
			ContainerPort: 15672,
			HostPort:      15672,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Default Username",
			Key:     "RABBITMQ_DEFAULT_USER",
			Value:   "guest",
			Mutable: true,
		},
		{
			Text:    "Default Password",
			Key:     "RABBITMQ_DEFAULT_PASS",
			Value:   "guest",
			Mutable: true,
		},
	},
}
