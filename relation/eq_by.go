package relation

// EqBy returns a function that checks if the yield results of the base and input values are equal.
// The type parameter T must be comparable.
func EqBy[T comparable](yield func(T) T, base T) func(T) bool {
	return func(val T) bool {
		return yield(base) == yield(val)
	}
}
