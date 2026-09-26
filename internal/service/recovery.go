package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wyw14/solo-036/internal/emergency"
	"github.com/wyw14/solo-036/internal/model"
	"github.com/wyw14/solo-036/internal/route"
)

func (r *Runtime) RecoveryBarrier() error {
	if err := ValidateRecoveryOrdering(); err != nil {
		return err
	}
	if !r.JournalHealthy() {
		return errors.New("journal is not healthy")
	}
	for _, lease := range route.Snapshot(r.Routes) {
		current, ok := r.Runs.Get(lease.RunID)
		if !ok {
			return fmt.Errorf("route %s references missing run %s", lease.Route, lease.RunID)
		}
		if current.State.Terminal() {
			return fmt.Errorf("route %s references terminal run %s", lease.Route, lease.RunID)
		}
	}
	for _, operation := range r.Operations() {
		if operation.State == model.StateDepressurizing && r.Pressure.Current() > 0 {
			incident := emergency.Incident(operation.ID, "pressure remains during depressurization")
			if emergency.IsCritical(incident) {
				return errors.New(incident.Message)
			}
		}
	}
	return nil
}

func (r *Runtime) WaitForSafeState(ctx context.Context, interval time.Duration) error {
	if interval <= 0 {
		interval = 10 * time.Millisecond
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		if err := r.RecoveryBarrier(); err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}
