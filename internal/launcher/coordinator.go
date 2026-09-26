package launcher

import (
	"fmt"
	"github.com/wyw14/solo-036/internal/model"
)

func (c *Controller) BeginClosure(runID string, epoch int64) {
	c.mu.Lock()
	c.current = model.CommandReceipt{RunID: runID, Command: "closure", Epoch: epoch, At: timeNow()}
	c.mu.Unlock()
}
func (c *Controller) ConfirmClosure(receipt model.CommandReceipt) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if receipt.RunID != c.current.RunID || receipt.Epoch != c.current.Epoch || !receipt.Accepted {
		return fmt.Errorf("closure proof rejected")
	}
	c.pressure.Authorize()
	return nil
}
func (c *Controller) CanLaunch(receipt model.CommandReceipt) bool {
	return c.permits.AllowReceive(receipt.RunID, receipt, receipt.Epoch) && c.pressure.CanPressurize(receipt)
}
func (c *Controller) Current() model.CommandReceipt {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.current
}
