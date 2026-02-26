# services/

One file per service. Each file exports a single `types.Service` variable.

## Adding a new service

Create a file like `services/myservice.go`:

```go
package services

import "github.com/fhsinchy/tent/types"

var MyService types.Service = types.Service{
    Name:  "myservice",                    // lowercase, used as CLI arg and registry key
    Image: "docker.io/org/image",          // full image reference
    Tag:   "latest",                       // default tag, user can override interactively
    Volumes: []types.VolumeMount{          // omit if no persistence needed
        {Text: "Server Data Volume", Name: "myservice-data", Dest: "/data"},
    },
    PortMappings: []types.PortMapping{     // at least one required (used in ContainerName)
        {Text: "Server Port", ContainerPort: 5000, HostPort: 5000},
    },
    Env: []types.EnvVar{                   // omit if none needed
        {Text: "Password", Key: "MY_PASSWORD", Value: "secret", Mutable: true},
    },
    InsecureEnv:  []types.EnvVar{...},     // nil = insecure not supported, empty = drop all auth
    InsecureInfo: "description of what insecure mode does",
    Command:      []string{"arg1", "arg2"},// omit if image entrypoint is sufficient
}
```

Then add it to the registry in `store/services.go`:

```go
"myservice": services.MyService,
```

## Field reference

- **Name**: Lowercase identifier. Must match the registry key in `store/services.go`
- **Image**: Full image path including registry (e.g., `docker.io/redis`). No tag here
- **Tag**: Default image tag. User can override in interactive mode
- **Volumes**: Named volumes for data persistence. `Text` is the interactive prompt label, `Name` is the volume name prefix (gets prefixed with container name at runtime), `Dest` is the mount path
- **PortMappings**: At least one is required — `ContainerName()` uses the first entry's HostPort. `Text` is the prompt label
- **Env**: Environment variables. `Mutable: true` means the user gets prompted for a value in interactive mode. `Mutable: false` means the value is fixed
- **InsecureEnv**: If `nil`, `--insecure` flag returns an error for this service. If set (even empty), these env vars replace `Env` when insecure mode is active. An empty slice effectively drops all auth env vars
- **InsecureInfo**: Human-readable description printed when insecure mode activates (e.g., "username: root, password: (empty)")
- **Command**: Overrides the image entrypoint. Omit to use the image default

## Current services (24)

Cassandra, ClickHouse, CouchDB, DynamoDB, Elasticsearch, InfluxDB, MailHog, Mailpit, MariaDB, Meilisearch, Memcached, MinIO, MongoDB, MSSQL, MySQL, Neo4j, OpenSearch, PostGIS, PostgreSQL, RabbitMQ, Redis, SurrealDB, Typesense, Valkey

## Insecure mode support

ClickHouse, MongoDB, MySQL, Neo4j, PostgreSQL, SurrealDB
