package gumshoe

// Finding represents actions for all types of findings
type Finding interface {
	// Use finding data to come up with more findings.
	// A good measure of the usefulness of a finding
	// is how many findings are produced by this method.
	Investigate() []Finding

	// Get the name of the struct that implements the
	// Finding interface.
	GetTypeName() string

	// Get the ID of the finding. The ID can be any string
	// that is equivalent for two equal findings of the
	// same underlying struct (and different otherwise).
	GetID() string

	// Pretty-print the finding's useful data to stdout.
	Display()
}
