package services

import (
	"github.com/fhsinchy/tent/types"
)

// CouchDB service holds necessary data for creating and running the CouchDB container.
var CouchDB types.Service = types.Service{
	Name:  "couchdb",
	Image: "docker.io/couchdb",
	Tag:   "latest",
	Volumes: []types.VolumeMount{
		{
			Text: "Server Data Volume",
			Name: "couchdb-data",
			Dest: "/opt/couchdb/data",
		},
	},
	PortMappings: []types.PortMapping{
		{
			Text:          "Server Port",
			ContainerPort: 5984,
			HostPort:      5984,
		},
	},
	Env: []types.EnvVar{
		{
			Text:    "Admin Username",
			Key:     "COUCHDB_USER",
			Value:   "admin",
			Mutable: true,
		},
		{
			Text:    "Admin Password",
			Key:     "COUCHDB_PASSWORD",
			Value:   "secret",
			Mutable: true,
		},
	},
}
