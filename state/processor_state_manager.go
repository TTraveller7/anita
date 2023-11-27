package state

import (
	"fmt"
)

type ProcessorStateManager struct {
	stateStores map[string]StateStore
}

// TODO: What is this doing?
func (m *ProcessorStateManager) RegisterKeyValueStateStore(name string) error {
	return nil
}

func (m *ProcessorStateManager) GetKeyValueStateStore(name string) (KeyValueStateStore, error) {
	s, ok := m.stateStores[name]
	if !ok {
		return nil, fmt.Errorf("no key value state store found with name %s", name)
	}
	switch v := s.(type) {
	case KeyValueStateStore:
		return v, nil
	default:
		return nil, fmt.Errorf("no key value state store found with name %s", name)
	}
}
