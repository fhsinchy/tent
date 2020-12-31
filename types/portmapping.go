package types

import "github.com/containers/podman/v2/pkg/specgen"

// PortMapping type describes a single port mapping for a container.
type PortMapping struct {
	Name    string
	Mapping specgen.PortMapping
}
