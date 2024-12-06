package helper

import "github.com/xuender/helper/types"

func Tap[T any](yield func(T)) types.Mapper[T, T] {
	return func(val T) T {
		yield(val)

		return val
	}
}
