package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// PHPMyAdmin service holds necessary data for creating and running the PHPMyAdmin container.
var PHPMyAdmin types.Service = types.Service{
	Name:   "phpmyadmin",
	Image:  "docker.io/phpmyadmin",
	Tag:    "latest",
	Volume: specgen.NamedVolume{},
	PortMapping: specgen.PortMapping{
		ContainerPort: 80,
		HostPort:      8080,
	},
	Env: map[string]string{
		"PMA_ARBITRARY": "1",
	},
	Prompts: map[string]bool{
		"tag":  true,
		"port": true,
	},
}
