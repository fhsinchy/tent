package store

import (
	"github.com/fhsinchy/tent/services"
	"github.com/fhsinchy/tent/types"
)

// registry holds a collection of all the available services.
var registry = map[string]types.Service{
	"mysql":         services.MySQL,
	"mariadb":       services.MariaDB,
	"postgres":      services.Postgres,
	"postgis":       services.PostGIS,
	"mongo":         services.Mongo,
	"redis":         services.Redis,
	"memcached":     services.Memcached,
	"mailhog":       services.MailHog,
	"elasticsearch": services.ElasticSearch,
	"meilisearch":   services.MeiliSearch,
	"dynamodb":      services.DynamoDB,
	"mssql":         services.MicrosoftSQLServer,
}

// GetService returns a deep copy of the named service, or false if not found.
func GetService(name string) (types.Service, bool) {
	s, ok := registry[name]
	if !ok {
		return s, false
	}
	s.Env = append([]types.EnvVar(nil), s.Env...)
	if s.InsecureEnv != nil {
		s.InsecureEnv = append([]types.EnvVar(nil), s.InsecureEnv...)
	}
	s.PortMappings = append([]types.PortMapping(nil), s.PortMappings...)
	s.Volumes = append([]types.VolumeMount(nil), s.Volumes...)
	s.Command = append([]string(nil), s.Command...)
	return s, true
}
