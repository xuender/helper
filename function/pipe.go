package function

import "github.com/xuender/rg/types"

// Pipe composes multiple functions into a single function, applying them sequentially.
func Pipe[T any](yields ...types.Mapper[T, T]) types.Mapper[T, T] {
	return func(val T) T {
		for _, yield := range yields {
			val = yield(val)
		}

		return val
	}
}
