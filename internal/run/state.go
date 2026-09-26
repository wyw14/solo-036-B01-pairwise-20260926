package run

import (
	"github.com/wyw14/solo-036/internal/model"
	"sync"
)

type Store struct {
	mu   sync.RWMutex
	runs map[string]model.Run
}

func NewStore() *Store { return &Store{runs: make(map[string]model.Run)} }
