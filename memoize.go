package helper

import (
	"github.com/xuender/helper/types"
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
