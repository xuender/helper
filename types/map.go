package types

// Mapper defines a function type that transforms a value.
type Mapper[T any] func(T) T
