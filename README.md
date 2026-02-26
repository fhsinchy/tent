# tent ![GitHub all releases](https://img.shields.io/github/downloads/fhsinchy/tent/total)

Tent runs development dependencies (databases, caches, message brokers, mail catchers) as pre-configured Podman containers. No Dockerfiles, no compose files.

```bash
tent start postgres -d       # running in seconds
tent start redis mongo -d    # multiple at once
tent stop --all              # done for the day
```

Inspired by [tighten/takeout](https://github.com/tighten/takeout).

**Tent is for local development only.** It skips TLS, binds to all interfaces, sets weak default passwords, and has no backup or monitoring. Do not run it on a production server.

## Available services

Run `tent services` to get this list from the CLI.

| Service | Default ports |
|---------|--------------|
| Cassandra | 9042 |
| ClickHouse | 8123, 9000 |
| CouchDB | 5984 |
| DynamoDB | 8000 |
| Elasticsearch | 9200 |
| InfluxDB | 8086 |
| MailHog | 1025, 8025 |
| Mailpit | 1025, 8025 |
| MariaDB | 3306 |
| Meilisearch | 7700 |
| Memcached | 11211 |
| MinIO | 9000, 9001 |
| MongoDB | 27017 |
| MSSQL | 1433 |
| MySQL | 3306 |
| Neo4j | 7474, 7687 |
| OpenSearch | 9200, 9600 |
| PostGIS | 5432 |
| PostgreSQL | 5432 |
| RabbitMQ | 5672, 15672 |
| Redis | 6379 |
| SurrealDB | 8000 |
| Typesense | 8108 |
| Valkey | 6379 |

Some services share default ports (Redis/Valkey on 6379, MinIO/ClickHouse on 9000, etc.). Tent prompts you to pick a different port when you start a service interactively.

## Dependencies

- Linux
- [Podman](https://podman.io/getting-started/installation) installed
- Podman system service running

If you have Podman installed, start the system service with:

```bash
# start the service
systemctl --user start podman.socket

# make it survive reboots
systemctl --user enable podman.socket
```

Tent talks to Podman through the user socket, so the `--user` flag matters here.

## Installation

Release binaries are statically linked and have no runtime dependencies beyond Podman itself. Grab the binary for your platform from the [releases page](https://github.com/fhsinchy/tent/releases/), then:

```bash
chmod +x ./tent
sudo mv ./tent /usr/local/bin
```

### Build from source

Building from source requires Go 1.23+ and a C compiler. You can build in two ways:

**Static build (no C library dependencies):**

```bash
git clone https://github.com/fhsinchy/tent.git ~/tent
cd ~/tent
CGO_ENABLED=0 go build -tags containers_image_openpgp -o bin/tent .
```

**Dynamic build (links against system libraries):**

This requires development headers for gpgme, btrfs, and device-mapper.

Fedora / RHEL / CentOS:

```bash
sudo dnf groupinstall "Development Tools" -y
sudo dnf install golang btrfs-progs-devel gpgme-devel device-mapper-devel -y
```

Debian / Ubuntu:

```bash
sudo apt install build-essential golang-go libbtrfs-dev libgpgme-dev libdevmapper-dev -y
```

Arch Linux:

```bash
sudo pacman -S base-devel go btrfs-progs gpgme device-mapper
```

openSUSE:

```bash
sudo zypper install -t pattern devel_basis
sudo zypper install go libbtrfs-devel gpgme-devel device-mapper-devel
```

Then build and install:

```bash
git clone https://github.com/fhsinchy/tent.git ~/tent
cd ~/tent
make install
```

## Usage

### Starting services

```bash
# interactive mode — prompts for tag, ports, credentials
tent start mysql

# skip all prompts, use defaults
tent start mysql --default
# or
tent start mysql -d

# start several at once
tent start redis mongo -d
```

### Insecure mode

Some services support `--insecure` to disable authentication entirely. Useful when you just want to poke at something locally without worrying about passwords.

```bash
tent start postgres --insecure -d    # trust auth, no password
tent start neo4j --insecure -d       # auth disabled
tent start clickhouse --insecure -d  # empty password
```

Services that support it: ClickHouse, MongoDB, MySQL, Neo4j, PostgreSQL, SurrealDB.

### Restart policies

```bash
tent start redis -d --restart always
tent start postgres -d --restart unless-stopped
tent start mysql -d --restart on-failure:5
```

### Stopping services

```bash
# stop a specific service (prompts if multiple instances are running)
tent stop mysql

# stop all instances of a service
tent stop mysql --all

# stop everything tent is running
tent stop --all
```

### Listing running services

```bash
tent list
```

Prints a table with container names, images, and all mapped ports.

### Shell completion

`tent start` and `tent stop` support tab completion for service names. Set up shell completions with:

```bash
# bash
tent completion bash > /etc/bash_completion.d/tent

# zsh
tent completion zsh > "${fpath[1]}/_tent"

# fish
tent completion fish > ~/.config/fish/completions/tent.fish
```

## Running multiple versions

Since everything is a container, you can run multiple versions of the same service on different ports.

```bash
tent start mysql -d                  # latest on port 3306
tent start mysql                     # pick tag 5.7 and port 3307
tent list
```

```
 CONTAINER                   IMAGE                    PORTS
 tent-mysql-5.7-3307         docker.io/mysql:5.7      3307->3306/tcp
 tent-mysql-latest-3306      docker.io/mysql:latest   3306->3306/tcp
```

## Container management

Tent containers are regular Podman containers. You can use `podman logs`, `podman inspect`, and anything else you normally would. Tent is meant to get containers running quickly, not to replace the Podman CLI.

Most services persist data in named volumes. Stopping a service removes the container but keeps the volume. You can manage volumes with `podman volume ls` and `podman volume rm` as usual.
