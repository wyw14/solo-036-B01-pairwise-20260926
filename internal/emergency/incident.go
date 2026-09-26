package emergency

import "github.com/wyw14/solo-036/internal/model"

func Incident(runID, message string) model.Incident {
	return model.Incident{ID: runID + "/incident", RunID: runID, Severity: "high", Message: message}
}
func IsCritical(i model.Incident) bool { return i.Severity == "high" || i.Severity == "critical" }
