package types

// PortMapping represents a single port mapping for a container.
type PortMapping struct {
	Text          string
	HostPort      uint16
	ContainerPort uint16
}
