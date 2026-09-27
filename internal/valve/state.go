package valve

func (c *Controller) BeginIsolation(runID string, epoch int64) {
	c.mu.Lock()
	delete(c.isolated, runID)
	c.mu.Unlock()
}

func (c *Controller) CompleteIsolation(runID string, epoch int64) {
	c.mu.Lock()
	c.isolated[runID] = epoch
	c.mu.Unlock()
}

func (c *Controller) Release(runID string) {
	c.mu.Lock()
	delete(c.isolated, runID)
	c.mu.Unlock()
}
