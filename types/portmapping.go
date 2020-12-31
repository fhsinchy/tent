package types

import "github.com/containers/podman/v2/pkg/specgen"

// PortMapping represents a single port mapping for a container.
type PortMapping struct {
	Text    string
	Mapping specgen.PortMapping
}
