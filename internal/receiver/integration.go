package receiver

func Ready(c *Controller, runID string) bool { return c.IsOpen(runID) }
