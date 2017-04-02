package lifecycle

import (
	"sync"
	"testing"

	"github.com/ml-tv/tv-api/src/core/primitives/models"
)

var _models = &savedModels{
	list: make(map[testing.TB]map[interface{}]bool),
}

// savedModels represents a list of models grouped by Test
// Since tests are run in parallel, we need to use mutexes
type savedModels struct {
	sync.Mutex
	list map[testing.TB]map[interface{}]bool
}

// Push adds a new model to the list
func (sm *savedModels) Push(t testing.TB, obj models.FullyDeletable) {
	_models.Lock()
	defer _models.Unlock()

	if _, ok := sm.list[t]; !ok {
		sm.list[t] = make(map[interface{}]bool, 0)
	}

	sm.list[t][obj] = true
}

// Push adds a new model to the list
func (sm *savedModels) Purge(t testing.TB) {
	sm.Lock()
	defer sm.Unlock()

	list, ok := sm.list[t]
	if !ok {
		return
	}

	for obj := range list {
		deletable, ok := obj.(models.FullyDeletable)
		if !ok {
			t.Fatalf("could not delete saved object")
		}

		if err := deletable.FullyDelete(); err != nil {
			t.Fatalf("could not delete saved object: %s", err)
		}
	}

	delete(sm.list, t)
}

// SaveModels saves a list of models that can be purged using PurgeModels()
func SaveModels(t testing.TB, models ...models.FullyDeletable) {
	for _, model := range models {
		_models.Push(t, model)
	}
}

// PurgeModels removes all models stored for the given test
func PurgeModels(t testing.TB) {
	_models.Purge(t)
}
