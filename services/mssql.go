package services

import (
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// MicrosoftSQLServer service holds necessary data for creating and running the MicrosoftSQLServer container.
var MicrosoftSQLServer types.Service = types.Service{
	Name:  "mssql",
	Image: "docker.io/mcr.microsoft.com/mssql/server",
	Tag:   "latest",
	PortMappings: []types.PortMapping{
		{
			Text: "Server Port",
			Mapping: specgen.PortMapping{
				ContainerPort: 1433,
				HostPort:      1433,
			},
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Server Root Password",
			Key:     "SA_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
		{
			Text:    "Accept EULA",
			Key:     "ACCEPT_EULA",
			Value:   "Y",
			Mutable: false,
		},
	},
}
