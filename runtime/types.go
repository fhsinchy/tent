package runtime

// ContainerInfo holds tent-owned container metadata, replacing the Podman entities.ListContainer type.
type ContainerInfo struct {
	ID    string
	Name  string
	Image string
	Ports []PortInfo
}

// PortInfo holds port mapping details for a container.
type PortInfo struct {
	HostPort      uint16
	ContainerPort uint16
	Protocol      string
}
