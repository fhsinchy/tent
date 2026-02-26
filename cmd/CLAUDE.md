# cmd/

Cobra CLI commands. Each file registers one command with `rootCmd` in its `init()` function.

## Files

- **root.go** — Root command, Viper config initialization, `rootCmd.Version` wired to `version` variable from `version.go`
- **version.go** — `tent version` subcommand. Defines the `version` variable (default `"development"`, overridden by ldflags at build time)
- **start.go** — `tent start <service...>`. Flags: `--default`/`-d`, `--insecure`, `--restart`/`-r`. Has `ValidArgs` from `store.ListServiceNames()` for tab completion. Contains `promptForService()` which interactively asks for tag, ports, mutable env vars, and volume names. Logic lives in `startServices(engine, args, isDefault, insecure, restartPolicy)`
- **stop.go** — `tent stop [service...]`. Flags: `--all`/`-a`. Also has `ValidArgs`. When multiple containers match a service, prompts user to pick one (unless `--all`). Logic lives in `stopServices(engine, args, isAll)`
- **list.go** — `tent list`. Prints a tabwriter table of running tent containers with all port mappings joined by commas. Logic lives in `listContainers(engine, io.Writer) error`
- **services.go** — `tent services`. Prints all available service names sorted alphabetically, one per line

## Conventions

- Cobra `Run` closures call `runtime.Connect()` then delegate to extracted functions (`startServices`, `stopServices`, `listContainers`) that accept a `runtime.ContainerEngine` interface. This keeps Podman coupling in the closure and lets tests use a mock
- Invalid service names print a message with a hint to run `tent services`
- Multiple services can be passed as positional args to `start` and `stop`
- The `version` variable in `version.go` is shared with `root.go` (same package) to set `rootCmd.Version`
