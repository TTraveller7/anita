package processor

type ProcessorNode struct {
	name        string
	processor   Processor
	stateStores map[string]bool
}

func newProcessorNode(name string, processor Processor, stateStores map[string]bool) *ProcessorNode {
	return &ProcessorNode{
		name:        name,
		processor:   processor,
		stateStores: stateStores,
	}
}

type ProcessorNodeBuilder struct {
	Builder
	name        string
	processor   Processor
	stateStores map[string]bool
}

func NewProcessorNodeBuilder(name string, processor Processor) Builder {
	return &ProcessorNodeBuilder{
		name:        name,
		processor:   processor,
		stateStores: make(map[string]bool, 0),
	}
}

func (b *ProcessorNodeBuilder) Build() *ProcessorNode {
	return newProcessorNode(b.name, b.processor, b.stateStores)
}

func (b *ProcessorNodeBuilder) Name() string {
	return b.name
}

func (b *ProcessorNodeBuilder) AddStateStore(stateStore string) {
	b.stateStores[stateStore] = true
}
