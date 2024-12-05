package function

import (
	"slices"

	"github.com/xuender/rg/types"
)

// Compose composes multiple functions into a single function, applying them in reverse order.
func Compose[T any](yields ...types.Mapper[T, T]) types.Mapper[T, T] {
	slices.Reverse(yields)

	return Pipe(yields...)
}
