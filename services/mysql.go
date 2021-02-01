package services

import (
	"github.com/fhsinchy/tent/types"
)

// MySQL service holds necessary data for creating and running the MySQL container.
var MySQL types.Service = types.Service{
	Name:  "mysql",
	Image: "docker.io/mysql",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "mysql-data",
			Dest: "/var/lib/mysql",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 3306,
			HostPort:      3306,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Server Root Password",
			Key:     "MYSQL_ROOT_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
	},
}
