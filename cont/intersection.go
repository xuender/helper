package cont

// Intersection returns a slice containing elements that are present in all the provided slices.
func Intersection[S ~[]E, E comparable](items ...S) S {
	if len(items) == 0 {
		return nil
	}

	nums := map[E]int{}
	keys := []E{}

	for _, slice := range items {
		set := NewSet[E]()
		for _, item := range slice {
			if set.Has(item) {
				continue
			}

			set.Add(item)

			if num, has := nums[item]; has {
				nums[item] = num + 1
			} else {
				keys = append(keys, item)
				nums[item] = 1
			}
		}
	}

	return Filter(keys, func(item E) bool { return nums[item] == len(items) })
}
