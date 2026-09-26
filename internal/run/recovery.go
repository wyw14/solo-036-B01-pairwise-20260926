package run

import "github.com/wyw14/solo-036/internal/model"

func Restore(store *Store, records []model.Run) {
	store.mu.Lock()
	defer store.mu.Unlock()
	for _, r := range records {
		if model.ValidateRun(r) {
			store.runs[r.ID] = r
		}
	}
}
func Active(store *Store) []model.Run {
	records := store.Snapshot()
	out := make([]model.Run, 0, len(records))
	for _, r := range records {
		if !r.State.Terminal() {
			out = append(out, r)
		}
	}
	return out
}
