package cont

import "github.com/xuender/helper/gtype"

func Find[S ~[]E, E any](slice S, checker gtype.Checker[E]) (E, bool) {
	for _, item := range slice {
		if checker(item) {
			return item, true
		}
	}

	var zero E

	return zero, false
}
