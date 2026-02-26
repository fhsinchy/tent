package services

import (
	"github.com/fhsinchy/tent/types"
)

// InfluxDB service holds necessary data for creating and running the InfluxDB container.
var InfluxDB types.Service = types.Service{
	Name:  "influxdb",
	Image: "docker.io/influxdb",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "influxdb-data",
			Dest: "/var/lib/influxdb2",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 8086,
			HostPort:      8086,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Init Mode",
			Key:     "DOCKER_INFLUXDB_INIT_MODE",
			Value:   "setup",
			Mutable: false,
		},
		{
			Text:    "Admin Username",
			Key:     "DOCKER_INFLUXDB_INIT_USERNAME",
			Value:   "admin",
			Mutable: true,
		},
		{
			Text:    "Admin Password",
			Key:     "DOCKER_INFLUXDB_INIT_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
		{
			Text:    "Organization",
			Key:     "DOCKER_INFLUXDB_INIT_ORG",
			Value:   "tent",
			Mutable: true,
		},
		{
			Text:    "Default Bucket",
			Key:     "DOCKER_INFLUXDB_INIT_BUCKET",
			Value:   "default",
			Mutable: true,
		},
	},
}
