package route

import "fmt"

func (c *Coordinator) Advance(runID string, position float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, lease := range c.owners {
		if lease.RunID == runID {
			if position < c.progress[runID] {
				return fmt.Errorf("position regressed")
			}
			c.progress[runID] = position
			return nil
		}
	}
	return fmt.Errorf("run has no route")
}
