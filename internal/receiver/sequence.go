package receiver

import "github.com/wyw14/solo-036/internal/model"

func Prepare(c *Controller, runID string, epoch int64) { c.Expect(runID, epoch) }
func Receive(c *Controller, runID string, proof model.CommandReceipt) bool {
	return c.Open(runID, proof) == nil
}
