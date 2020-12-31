package store

// Envs map holds the information for nevironment variables.
var Envs = map[string]map[string]string{
	"mongo": {
		"username": "MONGO_INITDB_ROOT_USERNAME",
		"password": "MONGO_INITDB_ROOT_PASSWORD",
	},
	"mysql": {
		"password": "MYSQL_ROOT_PASSWORD",
	},
	"mariadb": {
		"password": "MYSQL_ROOT_PASSWORD",
	},
	"postgres": {
		"password": "POSTGRES_PASSWORD",
	},
}
