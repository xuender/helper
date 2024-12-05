package list_test

import (
	"testing"

	"github.com/xuender/rg/list"
	"github.com/xuender/rg/math"
)

func BenchmarkReduce(b *testing.B) {
	adder := list.Reduce(func(val1, val2 int) int {
		return val1 + val2
	})

	b.ResetTimer()

	for range b.N {
		adder([]int{1, 2, 3, 4})
	}
}

func BenchmarkReduceCurry(b *testing.B) {
	adder := list.ReduceCurry(math.Add[int])

	b.ResetTimer()

	for range b.N {
		adder([]int{1, 2, 3, 4})
	}
}
