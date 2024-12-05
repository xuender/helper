package function

import "github.com/xuender/rg/types"

func Tap[T any](yield func(T)) types.Mapper[T, T] {
	return func(val T) T {
		yield(val)

		return val
	}
}
