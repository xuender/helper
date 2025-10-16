package cont

import "github.com/xuender/helper/types"

// Negate inverts the sign of each numeric element in the slice
//
// Parameters:
//   - items: a slice implementing ~[]E interface, where E must be a numeric type
//
// Returns:
//   - S: the slice with inverted signs, same type as input slice
//
// Generic constraints:
//   - S: must be a slice type with elements of type E
//   - E: must be a numeric type (integer, float, etc.)
func Negate[S ~[]E, E types.Number](items S) S {
	for index := range items {
		items[index] = -items[index]
	}

	return items
}
