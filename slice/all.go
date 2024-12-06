package slice

import "github.com/xuender/helper/types"

func All[T any](items []T, chekders ...types.Checker[T]) bool {
	for _, checker := range chekders {
		for _, item := range items {
			if !checker(item) {
				return false
			}
		}
	}

	return true
}
