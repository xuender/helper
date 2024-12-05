package function_test

import (
	"testing"

	"github.com/xuender/rg/function"
)

func BenchmarkCurryReduce(b *testing.B) {
	adder := function.CurryReducer(func(val1, val2 int) int { return val1 + val2 })
	addOne := adder(1)
	addTwo := adder(2)

	b.ResetTimer()

	for val := range b.N {
		addOne(val)
		addTwo(val)
	}
}

func BenchmarkAddOne(b *testing.B) {
	addOne := func(val int) int { return val + 1 }
	addTwo := func(val int) int { return val + 2 }

	b.ResetTimer()

	for val := range b.N {
		addOne(val)
		addTwo(val)
	}
}
