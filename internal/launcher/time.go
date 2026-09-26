package launcher

import "time"

func timeNow() time.Time       { return time.Now().UTC() }
func Ready(c *Controller) bool { return c != nil && c.pressure != nil && c.permits != nil }
