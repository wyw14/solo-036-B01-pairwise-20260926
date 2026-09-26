package journal

import (
	"github.com/wyw14/solo-036/internal/model"
)

func (s *Store) Append(e model.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.Durable = true
	s.events = append(s.events, e)
	return nil
}
func (s *Store) Events() []model.Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]model.Event(nil), s.events...)
}
