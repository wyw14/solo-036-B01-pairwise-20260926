package tracker

import "github.com/wyw14/solo-036/internal/model"

func (c *Coordinator) ArrivalQuorum(events []model.PassageEvidence, required int) bool {
	sources := make(map[string]bool)
	for _, e := range events {
		if model.ValidEvidence(e) {
			sources[e.SourceID] = true
		}
	}
	return len(sources) >= required
}
