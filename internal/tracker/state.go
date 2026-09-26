package tracker

import (
	"github.com/wyw14/solo-036/internal/model"
	"sync"
)

type Coordinator struct {
	mu         sync.Mutex
	points     map[string][]model.PassageEvidence
	generation map[string]int64
	sources    map[string]map[string]bool
}

func NewCoordinator() *Coordinator {
	return &Coordinator{points: make(map[string][]model.PassageEvidence), generation: make(map[string]int64), sources: make(map[string]map[string]bool)}
}
