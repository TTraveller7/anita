package anita

import (
	"fmt"
	"ttraveller7/anita/config"
	"ttraveller7/anita/processor"
	"ttraveller7/anita/state"
)

type Topology struct {
	Conf               *config.AnitaConfig
	stateStoreBuilders map[string]state.Builder
	processorBuilders  map[string]processor.Builder
}

func (t *Topology) AddStateStore(builder state.Builder, processorNames ...string) error {
	builderName := builder.Name()
	if builderName == "" {
		return fmt.Errorf("builder name cannot be nil")
	}
	if _, exists := t.stateStoreBuilders[builderName]; exists {
		return fmt.Errorf("another state store builder already has name %s", builderName)
	}
	t.stateStoreBuilders[builderName] = builder

	// connect processor with state store
	for _, processorName := range processorNames {
		processorBuilder, exists := t.processorBuilders[processorName]
		if !exists {
			return fmt.Errorf("processor %s is not added yet", processorName)
		}
		processorBuilder.AddStateStore(builder.Name())
	}

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

func (t *Topology) NumOfRoutines() int {
	return len(t.processorBuilders)
}

func NewTopology(conf *config.AnitaConfig) *Topology {
	c := conf
	if c == nil {
		// TODO: default config
		c = &config.AnitaConfig{}
	}
	return &Topology{
		Conf: c,
	}
}
