package relation

import "cmp"

func Gt[T cmp.Ordered](base T) func(T) bool {
	return func(data T) bool {
		return base > data
	}
}
