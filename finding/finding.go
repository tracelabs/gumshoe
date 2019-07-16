package finding

// Finding represents actions for all types of findings
type Finding interface {
	Investigate() []Finding
	GetUID() string
	GetType() string
}
