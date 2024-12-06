package helper

import "github.com/xuender/helper/types"

// Pipe composes multiple functions into a single function, applying them sequentially.
func Pipe[T any](mapper ...types.Mapper[T, T]) types.Mapper[T, T] {
	return func(val T) T {
		for _, yield := range mapper {
			val = yield(val)
		}

		return val
	}
}
