package anita

import "context"

type KafkaStreams struct {
	routines []*StreamRoutine
}

func (k *KafkaStreams) Start(ctx context.Context) error {
	for _, routine := range k.routines {
		if err := routine.start(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (k *KafkaStreams) Close() error {
	return nil
}
