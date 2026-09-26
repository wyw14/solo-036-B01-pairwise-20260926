package pressure

func (c *Controller) Recover(value float64) {
	c.mu.Lock()
	c.current = value
	c.enabled = false
	c.stopped = false
	c.mu.Unlock()
}
func (c *Controller) Reset() {
	c.mu.Lock()
	c.current = 0
	c.history = nil
	c.enabled = false
	c.stopped = false
	c.mu.Unlock()
}
