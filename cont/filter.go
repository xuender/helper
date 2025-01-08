package cont

import (
	"github.com/xuender/helper/types"
)

// Filter filters elements in a slice based on a provided checker function.
// It returns a new slice containing only the elements that satisfy the checker.
func Filter[S ~[]T, T any](items S, checker types.Checker[T]) S {
	ret := make([]T, 0, len(items))

	for _, item := range items {
		if checker(item) {
			ret = append(ret, item)
		}
	}

	return ret[:len(ret):len(ret)]
}
