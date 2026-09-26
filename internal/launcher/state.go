package launcher

import (
	"github.com/wyw14/solo-036/internal/model"
	"github.com/wyw14/solo-036/internal/permit"
	"github.com/wyw14/solo-036/internal/pressure"
	"sync"
)

type Controller struct {
	mu       sync.Mutex
	current  model.CommandReceipt
	pressure *pressure.Controller
	permits  *permit.Evaluator
}

func NewController(p *pressure.Controller, e *permit.Evaluator) *Controller {
	return &Controller{pressure: p, permits: e}
}
