# store/

Service registry. Single file: `services.go`.

## What it does

Maps string service names to `types.Service` values. Two public functions:

- **GetService(name string) (types.Service, bool)** — Returns a deep copy of the service by name. Deep copy is important: callers (especially `start` command) mutate the returned service during interactive prompts, and without the copy, those mutations would corrupt the registry for subsequent calls.
- **ListServiceNames() []string** — Returns all registered service names, sorted alphabetically. Used by `tent services` command and `ValidArgs` on both `start` and `stop` commands for tab completion.

## When you add a service

Add one line to the `registry` map:

```go
"servicename": services.ServiceVar,
```

The name must be lowercase and match `Service.Name`. Everything else (CLI availability, tab completion, `tent services` output) follows automatically.
