package anita

import (
	"fmt"
	"ttraveller7/anita/processor"
	"ttraveller7/anita/state"
)

type Topology struct {
	stateStoreBuilders map[string]state.Builder
	processorBuilders  map[string]processor.Builder
}

func (t *Topology) AddStateStore(builder state.Builder) error {
	builderName := builder.Name()
	if builderName == "" {
		return fmt.Errorf("builder name cannot be nil")
	}
	if _, exists := t.stateStoreBuilders[builderName]; exists {
		return fmt.Errorf("another state store builder already has name %s", builderName)
	}
	t.stateStoreBuilders[builderName] = builder
	return nil
}

func (t *Topology) AddProcessor(builder processor.Builder) error {
	name := builder.Name()
	if _, exists := t.processorBuilders[name]; exists {
		return fmt.Errorf("another processor builder already has name %s", name)
	}
	t.processorBuilders[name] = builder
	return nil
}

func NewTopology() *Topology {
	return &Topology{}
}
