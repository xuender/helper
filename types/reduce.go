package types

// Reducer defines a function type that reduces two values to one.
type Reducer[T any] func(val1, val2 T) T

// ReducerSlice defines a function type that reduces a slice of values to one.
type ReducerSlice[T any] func(items []T) T
