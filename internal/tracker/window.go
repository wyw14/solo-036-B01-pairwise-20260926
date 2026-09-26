package tracker

import "time"

func (c *Coordinator) Clearance(runID, station string, maxGap time.Duration) bool {
	points := c.Evidence(runID, station)
	if len(points) < 2 {
		return false
	}
	for i := 1; i < len(points); i++ {
		if points[i].At.Sub(points[i-1].At) > maxGap {
			return false
		}
		if points[i].Generation != points[0].Generation {
			return false
		}
	}
	return true
}
