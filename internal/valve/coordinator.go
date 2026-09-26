package valve

import (
	"github.com/wyw14/solo-036/internal/pressure"
	"sync"
)

type Controller struct {
	mu       sync.Mutex
	isolated map[string]int64
	pressure *pressure.Controller
}

func NewController(p *pressure.Controller) *Controller {
	return &Controller{isolated: make(map[string]int64), pressure: p}
}
