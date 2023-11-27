package anita

import (
	"context"
	"fmt"
)

type StreamRoutine struct {
}

func (r *StreamRoutine) start(ctx context.Context) (err error) {
	defer func() {
		if pErr := recover(); pErr != nil {
			err = fmt.Errorf("%v", pErr)
		}
	}()
	go r.runLoop(ctx)
	return nil
}

func (r *StreamRoutine) runLoop(ctx context.Context) {

}
