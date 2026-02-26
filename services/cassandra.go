package services

import (
	"github.com/fhsinchy/tent/types"
)

// Cassandra service holds necessary data for creating and running the Cassandra container.
var Cassandra types.Service = types.Service{
	Name:  "cassandra",
	Image: "docker.io/cassandra",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "cassandra-data",
			Dest: "/var/lib/cassandra",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 9042,
			HostPort:      9042,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Cluster Name",
			Key:     "CASSANDRA_CLUSTER_NAME",
			Value:   "tent-cluster",
			Mutable: true,
		},
	},
}
