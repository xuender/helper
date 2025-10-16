package gtype

// Reducer defines a function type that reduces two values to one.
type Reducer[T any] func(val1, val2 T) T
