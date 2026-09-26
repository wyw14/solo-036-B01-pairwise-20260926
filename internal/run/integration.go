package run

import "github.com/wyw14/solo-036/internal/model"

func Start(store *Store, id string) error {
	if err := store.Transition(id, model.StatePressurizing); err != nil {
		return err
	}
	return store.Transition(id, model.StateLaunched)
}

func Track(store *Store, id string) error { return store.Transition(id, model.StateTracking) }
