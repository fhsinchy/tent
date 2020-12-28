package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

var MariaDB types.Service = types.Service{
	Name:      "mariadb",
	Container: "tent-mariadb",
	Image:     "docker.io/mariadb",
	Tag:       "latest",
	Volume: specgen.NamedVolume{
		Name: "tent-mariadb-data",
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
}
