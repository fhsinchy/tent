package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// MySQL service holds necessary data for creating and running the MySQL container.
var MySQL types.Service = types.Service{
	Name:  "mysql",
	Image: "docker.io/mysql",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/var/lib/mysql",
	},
	PortMappings: []types.PortMapping{
		{
			Name: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 3306,
				HostPort:      3306,
			},
		},
	},
	Env: map[string]string{
		"MYSQL_ROOT_PASSWORD": "secret",
	},
	HasVolumes: true,
	Prompts: map[string]bool{
		"tag":      true,
		"password": true,
		"volume":   true,
		"port":     true,
	},
}
