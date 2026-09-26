package route

import "github.com/wyw14/solo-036/internal/model"

func Restore(c *Coordinator, leases []model.RouteLease) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, lease := range leases {
		if lease.Route != "" && lease.RunID != "" {
			c.owners[lease.Route] = lease
		}
	}
}
func Snapshot(c *Coordinator) []model.RouteLease {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]model.RouteLease, 0, len(c.owners))
	for _, lease := range c.owners {
		out = append(out, lease)
	}
	return out
}
