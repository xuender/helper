package slice

import "github.com/xuender/helper/types"

func First[T any](items []T, checker types.Checker[T]) (T, bool) {
	for _, item := range items {
		if checker(item) {
			return item, true
		}
	}

	return *new(T), false
}
