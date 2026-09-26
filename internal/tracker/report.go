package tracker

func Ready(c *Coordinator) bool { return c != nil && c.points != nil && c.generation != nil }
