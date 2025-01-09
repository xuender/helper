package cont

func Fill[S ~[]E, E any](slice S, val E, startAndEnd ...int) {
	end := len(slice)
	if end == 0 {
		return
	}

	start := 0

	switch len(startAndEnd) {
	case 0:
	case 1:
		start = startAndEnd[0]
	default:
		start = startAndEnd[0]
		end = startAndEnd[1]
	}

	for idx := start; idx < end; idx++ {
		slice[idx] = val
	}
}
