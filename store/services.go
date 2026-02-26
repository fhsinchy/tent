package store

import (
	"sort"

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
	"rabbitmq":      services.RabbitMQ,
	"minio":         services.MinIO,
	"clickhouse":    services.ClickHouse,
	"mailpit":       services.Mailpit,
	"couchdb":       services.CouchDB,
	"cassandra":     services.Cassandra,
	"neo4j":         services.Neo4j,
	"influxdb":      services.InfluxDB,
	"typesense":     services.Typesense,
	"surrealdb":     services.SurrealDB,
	"valkey":        services.Valkey,
	"opensearch":    services.OpenSearch,
}

// GetService returns a deep copy of the named service, or false if not found.
func GetService(name string) (types.Service, bool) {
	s, ok := registry[name]
	if !ok {
		return s, false
	}
	s.Env = append([]types.EnvVar(nil), s.Env...)
	if s.InsecureEnv != nil {
		cp := make([]types.EnvVar, len(s.InsecureEnv))
		copy(cp, s.InsecureEnv)
		s.InsecureEnv = cp
	}
	s.PortMappings = append([]types.PortMapping(nil), s.PortMappings...)
	s.Volumes = append([]types.VolumeMount(nil), s.Volumes...)
	s.Command = append([]string(nil), s.Command...)
	return s, true
}

// ListServiceNames returns a sorted list of all available service names.
func ListServiceNames() []string {
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
