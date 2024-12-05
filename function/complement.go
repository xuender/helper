package function

import "github.com/xuender/rg/types"

func Complement[T any](checker types.Checker[T]) types.Checker[T] {
	return func(val T) bool {
		return !checker(val)
	}
}
