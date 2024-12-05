package relation

// Equals returns a function that checks if the input value is equal to the given base value.
// The type parameter T must implement the comparable interface.
func Equals[T comparable](base T) func(T) bool {
	return func(val T) bool {
		return val == base
	}
}
