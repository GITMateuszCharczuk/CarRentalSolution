package base

import "seeder-manager/reference_store"

type Seeder interface {
	Seed(store *reference_store.InMemoryStore, token string) error
	Cleanup() error
}
