package types

// EnvVar represents a single environement variable for a container.
type EnvVar struct {
	Name    string
	Key     string
	Value   string
	Mutable bool
}
