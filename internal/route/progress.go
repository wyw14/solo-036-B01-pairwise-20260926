package route

func (c *Coordinator) Position(runID string) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.progress[runID]
}
