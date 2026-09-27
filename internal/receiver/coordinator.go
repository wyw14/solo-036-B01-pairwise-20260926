package receiver

import (
	"fmt"
	"github.com/wyw14/solo-036/internal/model"
)

func (c *Controller) Expect(runID string, epoch int64) {
	c.mu.Lock()
	c.current[runID] = epoch
	c.mu.Unlock()
}
func (c *Controller) Open(runID string, proof model.CommandReceipt) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	expected, ok := c.current[runID]
	if !ok || proof.RunID != runID || proof.Epoch != expected || !proof.Accepted {
		return fmt.Errorf("current isolation proof required")
	}
	c.opened[runID] = true
	return nil
}
func (c *Controller) IsOpen(runID string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.opened[runID]
}
func (c *Controller) Release(runID string) {
	c.mu.Lock()
	delete(c.current, runID)
	delete(c.opened, runID)
	c.mu.Unlock()
}
