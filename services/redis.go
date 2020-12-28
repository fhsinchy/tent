package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

var Redis types.Service = types.Service{
	Name:      "redis",
	Container: "tent-redis",
	Image:     "docker.io/redis",
	Tag:       "latest",
	Volume: specgen.NamedVolume{
		Name: "tent-redis-data",
		Dest: "/data",
	},
	PortMapping: specgen.PortMapping{
		ContainerPort: 6379,
		HostPort:      6379,
	},
	Env:        map[string]string{},
	HasVolumes: true,
}
