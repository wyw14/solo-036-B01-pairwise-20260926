package run

func Health(store *Store) bool { return store != nil && store.runs != nil }
