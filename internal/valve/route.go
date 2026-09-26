package valve

import "github.com/wyw14/solo-036/internal/model"

func Command(runID string, epoch int64, accepted bool) model.CommandReceipt {
	return model.CommandReceipt{RunID: runID, Command: "valve", Epoch: epoch, Accepted: accepted}
}
func Active(c *Controller, runID string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.isolated[runID]
	return ok
}
