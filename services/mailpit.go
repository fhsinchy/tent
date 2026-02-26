package services

import (
	"github.com/fhsinchy/tent/types"
)

// Mailpit service holds necessary data for creating and running the Mailpit container.
var Mailpit types.Service = types.Service{
	Name:  "mailpit",
	Image: "docker.io/axllent/mailpit",
	Tag:   "latest",
	PortMappings: []types.PortMapping{
		{
			Text:          "SMTP Port",
			ContainerPort: 1025,
			HostPort:      1025,
		},
		{
			Text:          "Web UI Port",
			ContainerPort: 8025,
			HostPort:      8025,
		},
	},
}
