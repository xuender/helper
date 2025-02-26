package fp_test

import (
	"testing"

	"github.com/xuender/helper/fp"
)

func BenchmarkBoth(b *testing.B) {
	leapYear := fp.Both(
		func(year int) bool { return year%4 == 0 },
		func(year int) bool { return year%100 != 0 },
	)

	b.ResetTimer()

	for year := range b.N {
		leapYear(year)
	}
}

func BenchmarkBothOne(b *testing.B) {
	leapYear := fp.Both(
		func(year int) bool { return year%4 == 0 && year%100 != 0 },
	)

	b.ResetTimer()

	for year := range b.N {
		leapYear(year)
	}
}

func BenchmarkFunc(b *testing.B) {
	leapYear := func(year int) bool {
		return year%4 == 0 && year%100 != 0
	}

	b.ResetTimer()

	for year := range b.N {
		leapYear(year)
	}
}
