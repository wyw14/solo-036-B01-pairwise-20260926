package journal

func Durable(s *Store, runID string) bool {
	for _, e := range s.Events() {
		if e.RunID == runID && e.Durable {
			return true
		}
	}
	return false
}
