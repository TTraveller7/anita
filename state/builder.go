package state

type Builder interface {
	Build() (StateStore, error)
	Name() string
}
