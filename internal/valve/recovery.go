package valve

func (c *Controller) Restore(runID string, epoch int64) { c.CompleteIsolation(runID, epoch) }
