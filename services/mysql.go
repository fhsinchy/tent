package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

var MySQL types.Service = types.Service{
	Name:      "mysql",
	Container: "tent-mysql",
	Image:     "docker.io/mysql",
	Tag:       "latest",
	Volume: specgen.NamedVolume{
		Name: "tent-mysql-data",
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
