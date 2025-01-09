package cont

import "slices"

// Union returns a slice containing all unique elements from the provided slices.
func Union[S ~[]E, E comparable](items ...S) S {
	if len(items) == 0 {
		return nil
	}

	set := NewSet[E]()
	union := []E{}

	for _, slice := range items {
		for _, item := range slice {
			if set.Has(item) {
				continue
			}

			set.Add(item)
			union = append(union, item)
		}
	}

	return union
}

// UnionFunc computes the union of multiple slices using a custom equality function.
// It returns a new slice containing unique elements from all input slices based on the provided comparator.
func UnionFunc[S ~[]E, E any](equal func(E, E) bool, items ...S) S {
	if len(items) == 0 {
		return nil
	}

	union := []E{}

	for _, slice := range items {
		for _, item := range slice {
			if slices.ContainsFunc(union, func(val E) bool {
				return equal(val, item)
			}) {
				continue
			}

			union = append(union, item)
		}
	}

	return union
}
