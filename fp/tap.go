package fp

import "github.com/xuender/helper/gtype"

// Tap applies the yield function to each value without altering it.
// It returns a Mapper that passes values through unchanged.
func Tap[T any](yield func(T)) gtype.Mapper[T, T] {
	return func(val T) T {
		yield(val)

		return val
	}
}
