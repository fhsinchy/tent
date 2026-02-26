package services

import (
	"github.com/fhsinchy/tent/types"
)

// ClickHouse service holds necessary data for creating and running the ClickHouse container.
var ClickHouse types.Service = types.Service{
	Name:  "clickhouse",
	Image: "docker.io/clickhouse/clickhouse-server",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "clickhouse-data",
			Dest: "/var/lib/clickhouse",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "HTTP Port",
			ContainerPort: 8123,
			HostPort:      8123,
		},
		{
			Text:          "Native Port",
			ContainerPort: 9000,
			HostPort:      9000,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Default Database",
			Key:     "CLICKHOUSE_DB",
			Value:   "default",
			Mutable: true,
		},
		{
			Text:    "Default Username",
			Key:     "CLICKHOUSE_USER",
			Value:   "default",
			Mutable: true,
		},
		{
			Text:    "Default Password",
			Key:     "CLICKHOUSE_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
	},
	InsecureEnv: []types.EnvVar{
		{
			Key:   "CLICKHOUSE_DB",
			Value: "default",
		},
		{
			Key:   "CLICKHOUSE_USER",
			Value: "default",
		},
	},
	InsecureInfo: "username: default, password: (empty)",
}
