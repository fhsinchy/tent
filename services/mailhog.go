package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// MailHog service holds necessary data for creating and running the MailHog container.
var MailHog types.Service = types.Service{
	Name:  "mailhog",
	Image: "docker.io/mailhog/mailhog",
	Tag:   "latest",
	PortMappings: []types.PortMapping{
		{
			Name: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 1025,
				HostPort:      1025,
			},
		},
		{
			Name: "Web UI Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 8025,
				HostPort:      8025,
			},
		},
	},
	Env:        map[string]string{},
	HasVolumes: false,
	Prompts: map[string]bool{
		"tag":  true,
		"port": true,
	},
}
