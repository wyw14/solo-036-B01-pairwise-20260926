package permit

import (
	"github.com/wyw14/solo-036/internal/journal"
	"github.com/wyw14/solo-036/internal/model"
	"time"
)

func EvaluateRepressurization(writer journal.Writer, evaluator *Evaluator, proof model.PassageEvidence) bool {
	event := model.Event{ID: proof.RunID + "/passage", RunID: proof.RunID, Kind: "passage", Epoch: proof.Generation, Payload: proof.Station, At: time.Now().UTC()}
	if err := writer.Append(event); err != nil {
		return false
	}
	return evaluator.AllowRepressurization(proof, true)
}
