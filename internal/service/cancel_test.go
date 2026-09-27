package service

import (
	"context"
	"testing"

	"github.com/wyw14/solo-036/internal/emergency"
	"github.com/wyw14/solo-036/internal/model"
)

func launchForCancelTest(t *testing.T, rt *Runtime) model.OperationRecord {
	t.Helper()
	op, err := rt.Launch(context.Background(), LaunchRequest{
		StationID:      "station-a",
		RouteID:        "main-station-a-station-c",
		TargetPressure: 8.5,
	})
	if err != nil {
		t.Fatalf("launch: %v", err)
	}
	return op
}

func TestCancelResetsFieldState(t *testing.T) {
	rt := NewRuntime()
	op := launchForCancelTest(t, rt)
	runID := op.Run.ID

	if err := rt.Cancel(runID); err != nil {
		t.Fatalf("cancel: %v", err)
	}

	snapshot := rt.Snapshot()
	record, ok := snapshot.Operation(runID)
	if !ok {
		t.Fatalf("cancelled run missing from snapshot")
	}
	if record.Run.State != model.StateCancelled {
		t.Errorf("run state = %s, want cancelled", record.Run.State)
	}

	// 压缩机停机、压力归零，现场处于安全态
	if !emergency.Safe(rt.Pressure) {
		t.Errorf("pressure controller current = %v, want 0 after cancel", rt.Pressure.Current())
	}
	ps, ok := lookupPressure(snapshot.Pressures, "station-a")
	if !ok {
		t.Fatalf("launcher pressure state missing")
	}
	if ps.RunID != "" || ps.Value != 0 || ps.Target != 0 || ps.Mode != model.CompressorStopped {
		t.Errorf("pressure state after cancel = %+v, want detached/zero/stopped", ps)
	}

	// 主阀关闭且不再归属该运行
	vs, ok := lookupValve(snapshot.Valves, "station-a/main")
	if !ok {
		t.Fatalf("main valve state missing")
	}
	if vs.Position != model.ValveClosed || vs.RunID != "" {
		t.Errorf("main valve after cancel = %+v, want closed/detached", vs)
	}

	// 关锁解除
	if _, locked := lookupClosure(snapshot.Closures, "station-a"); locked {
		t.Errorf("closure still present after cancel: %+v", snapshot.Closures)
	}

	// 路由租约释放
	rs, ok := lookupRoute(snapshot.Routes, "main-station-a-station-c")
	if !ok {
		t.Fatalf("route state missing")
	}
	if rs.Reserved || rs.RunID != "" || rs.LeaseEpoch != 0 {
		t.Errorf("route after cancel = %+v, want released", rs)
	}

	// 控制器围栏凭证已清除
	if rt.Launcher.Current().RunID != "" {
		t.Errorf("launcher closure fence still held: %+v", rt.Launcher.Current())
	}
	if rt.Receiver.IsOpen(runID) {
		t.Errorf("receiver still open for cancelled run")
	}
	if proof := rt.Valves.IsolationProof(runID, 0); proof.Accepted {
		t.Errorf("valve isolation fence still held for cancelled run")
	}

	// 取消后整个快照必须自洽可校验
	if err := snapshot.Validate(); err != nil {
		t.Errorf("snapshot invalid after cancel: %v", err)
	}
}

func TestCancelRecordsIncidentAndEvent(t *testing.T) {
	rt := NewRuntime()
	op := launchForCancelTest(t, rt)
	runID := op.Run.ID

	if err := rt.Cancel(runID); err != nil {
		t.Fatalf("cancel: %v", err)
	}

	foundIncident := false
	for _, incident := range rt.Incidents() {
		if incident.RunID == runID {
			foundIncident = true
		}
	}
	if !foundIncident {
		t.Errorf("expected an incident for cancelled run %s, got %+v", runID, rt.Incidents())
	}

	foundEvent := false
	for _, event := range rt.Journal.Events() {
		if event.RunID == runID && event.Kind == "cancelled" {
			foundEvent = true
			if event.Epoch <= 0 {
				t.Errorf("cancelled event missing epoch: %+v", event)
			}
		}
	}
	if !foundEvent {
		t.Errorf("expected cancelled journal event for run %s", runID)
	}
}

func TestCancelIdempotentForTerminalRun(t *testing.T) {
	rt := NewRuntime()
	op := launchForCancelTest(t, rt)
	if err := rt.Cancel(op.Run.ID); err != nil {
		t.Fatalf("first cancel: %v", err)
	}
	if err := rt.Cancel(op.Run.ID); err != nil {
		t.Errorf("second cancel should be idempotent, got %v", err)
	}
}

func lookupPressure(states []model.PressureState, station string) (model.PressureState, bool) {
	for _, s := range states {
		if s.StationID == station {
			return s, true
		}
	}
	return model.PressureState{}, false
}
func lookupValve(states []model.ValveState, id string) (model.ValveState, bool) {
	for _, s := range states {
		if s.ID == id {
			return s, true
		}
	}
	return model.ValveState{}, false
}
func lookupClosure(states []model.ClosureState, station string) (model.ClosureState, bool) {
	for _, s := range states {
		if s.StationID == station {
			return s, true
		}
	}
	return model.ClosureState{}, false
}
func lookupRoute(states []model.RouteState, id string) (model.RouteState, bool) {
	for _, s := range states {
		if s.ID == id {
			return s, true
		}
	}
	return model.RouteState{}, false
}
