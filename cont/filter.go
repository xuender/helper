package cont

import (
	"iter"

	"github.com/xuender/helper/gtype"
)

// Filter filters elements in a slice based on a provided checker function.
// It returns a new slice containing only the elements that satisfy the checker.
func Filter[S ~[]T, T any](items S, checker gtype.Checker[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if checker(item) {
				if !yield(item) {
					return
				}
			}
		}
	}
}
