package service

import (
	"context"
	"testing"

	"github.com/wyw14/solo-036/internal/model"
)

func launchForCancel(t *testing.T) *Runtime {
	t.Helper()
	runtime := NewRuntime()
	_, err := runtime.Launch(context.Background(), LaunchRequest{
		StationID:      "station-a",
		RouteID:        "main-station-a-station-c",
		TargetPressure: 6.5,
	})
	if err != nil {
		t.Fatalf("launch: %v", err)
	}
	return runtime
}

func activeRunID(t *testing.T, runtime *Runtime) string {
	t.Helper()
	active := runtime.Operations()
	if len(active) != 1 {
		t.Fatalf("expected one active run, got %d", len(active))
	}
	return active[0].ID
}

func TestCancelReleasesFieldState(t *testing.T) {
	runtime := launchForCancel(t)
	runID := activeRunID(t, runtime)

	if err := runtime.Cancel(runID); err != nil {
		t.Fatalf("cancel: %v", err)
	}

	snapshot := runtime.Snapshot()

	var cancelled bool
	for _, r := range snapshot.Runs {
		if r.ID == runID {
			cancelled = r.State == model.StateCancelled
		}
	}
	if !cancelled {
		t.Fatalf("run %s is not cancelled", runID)
	}

	for _, pressure := range snapshot.Pressures {
		if pressure.RunID == runID {
			t.Errorf("pressure %s still owned by cancelled run", pressure.StationID)
		}
		if pressure.Value != 0 || pressure.Target != 0 {
			t.Errorf("pressure %s not depressurized after cancel: %+v", pressure.StationID, pressure)
		}
		if pressure.StationID == "station-a" && pressure.Mode != model.CompressorStopped {
			t.Errorf("launcher compressor not stopped after cancel: %s", pressure.Mode)
		}
	}

	for _, valve := range snapshot.Valves {
		if valve.RunID == runID {
			t.Errorf("valve %s still owned by cancelled run", valve.ID)
		}
		if valve.Position != model.ValveClosed {
			t.Errorf("valve %s not closed after cancel: %s", valve.ID, valve.Position)
		}
	}

	for _, closure := range snapshot.Closures {
		if closure.RunID == runID {
			t.Errorf("closure %s still owned by cancelled run", closure.StationID)
		}
		if closure.Locked {
			t.Errorf("closure %s still locked after cancel", closure.StationID)
		}
	}

	for _, routeState := range snapshot.Routes {
		if routeState.Reserved || routeState.RunID != "" {
			t.Errorf("route %s still reserved after cancel: %+v", routeState.ID, routeState)
		}
	}

	if got := runtime.Pressure.Current(); got != 0 {
		t.Errorf("pressure controller still holds %.2f after cancel", got)
	}
	if err := snapshot.Validate(); err != nil {
		t.Errorf("snapshot invalid after cancel: %v", err)
	}
}

func TestCancelFieldStateMatchesJournalEpoch(t *testing.T) {
	runtime := launchForCancel(t)
	runID := activeRunID(t, runtime)

	if err := runtime.Cancel(runID); err != nil {
		t.Fatalf("cancel: %v", err)
	}

	var cancelEpoch int64 = -1
	for _, event := range runtime.Journal.Events() {
		if event.RunID == runID && event.Kind == "cancelled" {
			cancelEpoch = event.Epoch
		}
	}
	if cancelEpoch < 0 {
		t.Fatal("cancelled event missing from journal")
	}

	snapshot := runtime.Snapshot()
	for _, valve := range snapshot.Valves {
		if valve.StationID == "station-a" && valve.CommandEpoch != cancelEpoch {
			t.Errorf("valve %s command epoch %d does not match cancel event epoch %d", valve.ID, valve.CommandEpoch, cancelEpoch)
		}
	}
	for _, closure := range snapshot.Closures {
		if closure.StationID == "station-a" && closure.CommandEpoch != cancelEpoch {
			t.Errorf("closure %s command epoch %d does not match cancel event epoch %d", closure.StationID, closure.CommandEpoch, cancelEpoch)
		}
	}
}
