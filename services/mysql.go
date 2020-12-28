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
	},
	PortMapping: specgen.PortMapping{
		ContainerPort: 3306,
		HostPort:      3306,
	},
	Env: make(map[string]string),
}
