package finding

type Finding interface {
	GetUID() string
	GetType() string
}
