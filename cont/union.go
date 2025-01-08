package cont

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
