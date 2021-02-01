package types

// VolumeMount represents a single volume mount for a container.
type VolumeMount struct {
	Text string
	Name string
	Dest string
}
