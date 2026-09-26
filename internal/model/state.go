package model

type RunState string

const (
	StateLoaded         RunState = "loaded"
	StatePressurizing   RunState = "pressurizing"
	StateLaunched       RunState = "launched"
	StateTracking       RunState = "tracking"
	StateReceiving      RunState = "receiving"
	StateDepressurizing RunState = "depressurizing"
	StateComplete       RunState = "complete"
	StateCancelled      RunState = "cancelled"
)

func (s RunState) Terminal() bool { return s == StateComplete || s == StateCancelled }
func ValidTransition(from, to RunState) bool {
	if from == StateLoaded && (to == StatePressurizing || to == StateCancelled) {
		return true
	}
	if from == StatePressurizing && (to == StateLaunched || to == StateCancelled) {
		return true
	}
	if from == StateLaunched && (to == StateTracking || to == StateCancelled) {
		return true
	}
	if from == StateTracking && (to == StateReceiving || to == StateCancelled) {
		return true
	}
	if from == StateReceiving && (to == StateDepressurizing || to == StateCancelled) {
		return true
	}
	if from == StateDepressurizing && (to == StateComplete || to == StateCancelled) {
		return true
	}
	return false
}
