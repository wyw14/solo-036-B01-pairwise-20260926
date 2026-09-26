package pressure

import "github.com/wyw14/solo-036/internal/model"

func (c *Controller) CanPressurize(receipt model.CommandReceipt) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return receipt.Accepted && receipt.Epoch > 0 && c.enabled && !c.stopped
}
