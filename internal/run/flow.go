package run

import (
	"fmt"
	"github.com/wyw14/solo-036/internal/model"
)

func Receive(store *Store, id string) error {
	if err := store.Transition(id, model.StateReceiving); err != nil {
		return err
	}
	return store.Transition(id, model.StateDepressurizing)
}
func Complete(store *Store, id string) error {
	if err := store.Transition(id, model.StateComplete); err != nil {
		return fmt.Errorf("complete: %w", err)
	}
	return nil
}
