package math

import "cmp"

func Add[T cmp.Ordered](base T) func(T) T {
	return func(val T) T {
		return base + val
	}
}
