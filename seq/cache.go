package seq

import "iter"

func Cache[T any](input iter.Seq[T], size int) iter.Seq[T] {
	cache := make(chan T, size)

	go func() {
		for item := range input {
			cache <- item
		}

		close(cache)
	}()

	return func(yield func(T) bool) {
		for item := range cache {
			if !yield(item) {
				return
			}
		}
	}
}
