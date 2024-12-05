package relation

import (
	"cmp"

	"github.com/xuender/rg/types"
)

// Gt returns a function that checks if the input value is greater than the given base value.
// The type parameter T must implement the cmp.Ordered interface.
func Gt[T cmp.Ordered](base T) types.Checker[T] {
	return func(val T) bool {
		return base > val
	}
}
