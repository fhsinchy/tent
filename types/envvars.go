package types

// EnvVar represents a single environement variable for a container.
type EnvVar struct {
	Text    string
	Key     string
	Value   string
	Mutable bool
}
