package store

import (
	"github.com/fhsinchy/tent/services"
	"github.com/fhsinchy/tent/types"
)

// Services map holds a collection of all the available services.
var Services = map[string]*types.Service{
	"mysql":         &services.MySQL,
	"mariadb":       &services.MariaDB,
	"postgres":      &services.Postgres,
	"postgis":       &services.PostGIS,
	"mongo":         &services.Mongo,
	"redis":         &services.Redis,
	"memcached":     &services.Memcached,
	"mailhog":       &services.MailHog,
	"elasticsearch": &services.ElasticSearch,
	"meilisearch":   &services.MeiliSearch,
	"dynamodb":      &services.DynamoDB,
}
