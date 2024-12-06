package helper

import "github.com/xuender/helper/types"

// Tap applies the yield function to each value without altering it.
// It returns a Mapper that passes values through unchanged.
func Tap[T any](yield func(T)) types.Mapper[T, T] {
	return func(val T) T {
		yield(val)

		return val
	}
}
