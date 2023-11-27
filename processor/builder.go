package processor

type Builder interface {
	Build() *ProcessorNode
	Name() string
	AddStateStore(stateStore string)
}
