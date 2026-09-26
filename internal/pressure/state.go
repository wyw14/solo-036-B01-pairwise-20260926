package pressure

func (c *Controller) Current() float64 { c.mu.Lock(); defer c.mu.Unlock(); return c.current }
func (c *Controller) History() []float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]float64(nil), c.history...)
}
