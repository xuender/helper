package cont

import (
	"slices"

	"github.com/xuender/helper/types"
)

// Difference returns a new slice containing elements that are in src but not in any of the items slices.
// It uses the comparable constraint for equality checks.
func Difference[S ~[]E, E comparable](src S, items ...S) S {
	if len(items) == 0 {
		return src
	}

	diff := make([]E, 0, len(src))

	for _, item := range src {
		if Any(items, func(slice S) bool { return slices.Contains(slice, item) }) {
			continue
		}

		diff = append(diff, item)
	}

	return diff[:len(diff):len(diff)]
}

// DifferenceFunc returns a new slice containing elements that are in src but not in any of the items slices.
// It uses a custom equality function for comparisons.
func DifferenceFunc[S ~[]E, E any](equal types.Equaler[E], src S, items ...S) S {
	if len(items) == 0 {
		return src
	}

	diff := make([]E, 0, len(src))

	for _, item := range src {
		if Any(items, func(slice S) bool { return Contains(equal, slice, item) }) {
			continue
		}

		diff = append(diff, item)
	}

	return diff[:len(diff):len(diff)]
}
