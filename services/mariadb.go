package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// MariaDB service holds necessary data for creating and running the MariaDB container.
var MariaDB types.Service = types.Service{
	Name:  "mariadb",
	Image: "docker.io/mariadb",
	Tag:   "latest",
	Volume: specgen.NamedVolume{
		Dest: "/var/lib/mysql",
	},
	PortMapping: specgen.PortMapping{
		ContainerPort: 3306,
		HostPort:      3306,
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
