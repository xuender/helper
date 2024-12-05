package function

import (
	"sync"

	"github.com/xuender/rg/types"
)

// Memoize creates a memoized version of a mapper function.
func Memoize[K comparable, V any](mapper types.Mapper[K, V]) types.Mapper[K, V] {
	cache := map[K]V{}

	return func(key K) V {
		if val, has := cache[key]; has {
			return val
		}

		val := mapper(key)
		cache[key] = val

		return val
	}
}

// MemoizeSync creates a goroutine-safe memoized version of a mapper function.
func MemoizeSync[K comparable, V any](mapper types.Mapper[K, V]) types.Mapper[K, V] {
	cache := map[K]V{}
	mutex := sync.RWMutex{}

	return func(key K) V {
		mutex.RLock()

		if val, has := cache[key]; has {
			mutex.RUnlock()

			return val
		}

		mutex.RUnlock()
		mutex.Lock()

		defer mutex.Unlock()

		if val, has := cache[key]; has {
			return val
		}

		val := mapper(key)
		cache[key] = val

		return val
	}
}
