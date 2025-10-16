package cont

import (
	"github.com/xuender/helper/gtype"
)

// Contains checks if a slice contains an element based on a custom equality function.
// It returns true if the specified item is found in the slice.
func Contains[S ~[]E, E any](equal gtype.Equaler[E], items S, item E) bool {
	if len(items) == 0 {
		return false
	}

	for _, it := range items {
		if equal(item, it) {
			return true
		}
	}

	return false
}
