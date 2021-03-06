package services

import (
	"github.com/fhsinchy/tent/types"
)

// MariaDB service holds necessary data for creating and running the MariaDB container.
var MariaDB types.Service = types.Service{
	Name:  "mariadb",
	Image: "docker.io/mariadb",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "mariadb-data",
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
