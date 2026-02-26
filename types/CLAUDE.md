# types/

Shared type definitions used across packages. No logic beyond `Service` methods.

## Types

### Service (`service.go`)

The main type. Holds everything needed to create and run a container for a given service.

Methods:
- `ContainerName() string` — Generates `tent-{Name}-{Tag}-{HostPort}` using the first port mapping's HostPort
- `ApplyInsecure() (string, error)` — Replaces `Env` with `InsecureEnv`. Returns error if `InsecureEnv` is nil. Returns `InsecureInfo` string on success
- `ImageName() string` — Returns `{Image}:{Tag}`

### EnvVar (`envvar.go`)

Single environment variable. `Text` is the prompt label for interactive mode. `Mutable` controls whether the user can change the value interactively.

### PortMapping (`portmapping.go`)

Single port mapping. `Text` is the prompt label. `HostPort` and `ContainerPort` are `uint16`.

### VolumeMount (`volumemount.go`)

Single named volume. `Text` is the prompt label. `Name` is the base volume name (gets prefixed at runtime). `Dest` is the container mount path.

## Notes

- At least one `PortMapping` is required per service — `ContainerName()` indexes into `PortMappings[0]` without bounds checking
- The `Service` struct is designed to be copied by value, but its slice fields (Env, InsecureEnv, PortMappings, Volumes, Command) share underlying arrays. This is why `store.GetService()` deep-copies all slices before returning
