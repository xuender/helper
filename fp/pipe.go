package fp

import "github.com/xuender/helper/gtype"

// Pipe composes multiple functions into a single function, applying them sequentially.
func Pipe[T any](mapper ...gtype.Mapper[T, T]) gtype.Mapper[T, T] {
	return func(val T) T {
		for _, yield := range mapper {
			val = yield(val)
		}

		return val
	}
}
