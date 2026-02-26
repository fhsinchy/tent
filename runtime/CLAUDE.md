# runtime/

Wraps Podman v5 API bindings. All container operations go through the `Runtime` struct.

## Files

- **runtime.go** — `Runtime` struct and `Connect()`. Connects to the Podman socket at `$XDG_RUNTIME_DIR/podman/podman.sock`. Assumes rootless Podman (user socket, not system socket).
- **containers.go** — Container lifecycle methods on `Runtime`: `CreateContainer`, `StartContainer`, `StopContainer`, `RemoveContainer`, `ListTentContainers`. Also has the standalone `FilterContainers` function.
- **types.go** — `ContainerInfo` and `PortInfo` structs. These are tent's own types that replace Podman's `entities.ListContainer` to keep Podman types from leaking into the rest of the codebase.

## Container lifecycle

1. **CreateContainer(service, restartPolicy)** — Checks if container already exists (returns empty string if running). Pulls image if missing. Creates a specgen spec with ports, env, volumes, command, restart policy, and tent labels. Returns the new container ID.
2. **StartContainer(id)** — Starts the container and waits until it reaches `Running` state.
3. **StopContainer(id)** — Sends SIGTERM if the container is running.
4. **RemoveContainer(id)** — Removes the container if it's stopped. Does not remove volumes.

## Labels

Every tent container gets two labels:
- `tent.managed=true` — used by `ListTentContainers` to filter
- `tent.service=<name>` — used by `FilterContainers` to narrow down by service name

## Volume naming

Volume names at runtime are `{containerName}-{volumeName}`, not just the volume name from the service definition. This means multiple instances of the same service get separate volumes.
