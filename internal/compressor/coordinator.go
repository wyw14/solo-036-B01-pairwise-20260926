package compressor

import (
	"context"
	"github.com/wyw14/solo-036/internal/pressure"
	"github.com/wyw14/solo-036/internal/run"
	"time"
)

type Worker struct {
	pressure *pressure.Controller
	runs     *run.Store
	done     chan struct{}
}

func NewWorker(p *pressure.Controller, r *run.Store) *Worker {
	return &Worker{pressure: p, runs: r, done: make(chan struct{})}
}
func (w *Worker) Ramp(ctx context.Context, runID string, steps int) {
	go func() {
		defer close(w.done)
		for i := 0; i < steps; i++ {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Millisecond):
			}
			r, ok := w.runs.Get(runID)
			if !ok || r.State == "cancelled" {
				return
			}
			w.pressure.Apply(runID, float64(i+1))
		}
	}()
}
