package cont

import (
	"github.com/xuender/helper/gtype"
)

// Mean calculates the mean of the elements in the slice.
func Mean[S ~[]E, E gtype.Number](items S) E {
	if len(items) == 0 {
		return 0
	}

	return Sum(items) / E(len(items))
}
