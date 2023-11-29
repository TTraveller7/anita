package anita

import (
	"context"
)

type KafkaStreams struct {
	routines []*StreamRoutine
	topology *Topology
}

func NewKafkaStreams(topo *Topology) *KafkaStreams {
	k := &KafkaStreams{
		topology: topo,
	}

	for i := 1; i <= topo.NumOfRoutines(); i++ {
		k.createAndAddStreamRoutine(topo, i)
	}

	return k
}

func (k *KafkaStreams) createAndAddStreamRoutine(topo *Topology, routineIdx int) {

}

func (k *KafkaStreams) Start(ctx context.Context) error {
	for _, routine := range k.routines {
		if err := routine.start(ctx); err != nil {
			return err
		}
	}
	return nil
}

// TODO
func (k *KafkaStreams) Close() error {
	return nil
}
