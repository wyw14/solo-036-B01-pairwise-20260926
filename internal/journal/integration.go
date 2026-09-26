package journal

import (
	"github.com/wyw14/solo-036/internal/model"
	"sync"
)

type Writer interface{ Append(model.Event) error }

type Store struct {
	mu     sync.Mutex
	events []model.Event
}

func NewStore() *Store { return &Store{events: make([]model.Event, 0)} }
