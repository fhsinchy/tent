package store

import (
	"github.com/fhsinchy/tent/services"
	"github.com/fhsinchy/tent/types"
)

// Services map holds a collection of all the available services.
var Services = map[string]*types.Service{
	"mysql":    &services.MySQL,
	"mariadb":  &services.MariaDB,
	"postgres": &services.Postgres,
	"mongo":    &services.Mongo,
	"redis":    &services.Redis,
}
