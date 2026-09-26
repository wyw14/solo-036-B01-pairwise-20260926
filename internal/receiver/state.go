package receiver

import "sync"

type Controller struct {
	mu      sync.Mutex
	current map[string]int64
	opened  map[string]bool
}

func NewController() *Controller {
	return &Controller{current: make(map[string]int64), opened: make(map[string]bool)}
}
