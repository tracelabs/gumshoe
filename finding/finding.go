package finding

// Finding represents actions for all types of findings
type Finding interface {
	GetID() string
	Investigate() []Finding
}
